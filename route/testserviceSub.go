package route

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/ie310mu/ie310go/common/convert"
	jsonutil "github.com/ie310mu/ie310go/common/json"
	"github.com/ie310mu/ie310go/common/logsagent"
)

//要用全局变量，避免创建多个实例
var systemStateNip = genSystemStateNotiferInProcess()

//GetSystemState ..
func (s TestService) GetSystemState(args *ServiceArgs) NotifierInProcess {
	return systemStateNip
}

//ReciveSystemState 用于测试http订阅时的回调
func (s TestService) ReciveSystemState(args *ServiceArgs) *SystemState {
	datajson := args.GetStringParam("datajson")
	if datajson == "" {
		return nil
	}
	r := &SystemState{}
	jsonutil.FromJSON(datajson, r)
	return r
}

//SystemState ..
type SystemState struct {
	Time  string `json:"time"`
	State string `json:"state"`
}

func genSystemStateNotiferInProcess() *SystemStateNotiferInProcess {
	r := &SystemStateNotiferInProcess{}
	r.chsLock = new(sync.Mutex)
	r.chs = make(map[string]chan interface{})
	return r
}

//SystemStateNotiferInProcess ..
type SystemStateNotiferInProcess struct {
	BaseNotiferInProcess
	run int32
}

//Name ..
func (s *SystemStateNotiferInProcess) Name() string {
	return "SystemStateNotiferInProcess"
}

//TryRun ..
func (s *SystemStateNotiferInProcess) TryRun() {
	if atomic.LoadInt32(&s.run) == 1 {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			logsagent.Error("SystemStateNotiferInProcess run error： ")
			logsagent.Error(err)
			atomic.StoreInt32(&s.run, 0)
		}
	}()

	atomic.StoreInt32(&s.run, 1)

	go func() {
		for {
			if atomic.LoadInt32(&s.run) == 0 {
				break
			}
			time.Sleep(3 * time.Second)
			data := &SystemState{Time: convert.DateTimeWithSecToStr(time.Now()), State: "normal"}
			s.notiferChs(data)
		}
	}()
}

//Stop ..
func (s *SystemStateNotiferInProcess) Stop() {
	atomic.CompareAndSwapInt32(&s.run, 1, 0)
}
