package route

import (
	"sync"

	"github.com/ie310mu/ie310go/common/logsagent"
)

//Server ..
type Server interface {
	Name() string
	Run()
	Stop()
	Stopch() <-chan bool
	RegisterService(sv IService)
}

// server map
var servers = make(map[string]Server)

//RegisterServer ..
func RegisterServer(srv Server) {
	name := srv.Name()
	if _, dup := servers[name]; dup {
		logsagent.Warn("server " + name + " is duplicate")
		return
	}

	servers[name] = srv
}

//Run ..
func Run() {
	if len(servers) == 0 {
		registerTestServers()
	}

	var wg sync.WaitGroup
	for _, srv := range servers {
		wg.Add(1)        //这个要放在主协程，不然子协程来不及Add
		go run(srv, &wg) //在协程中启动，不影响其他server的启动
	}

	wg.Wait()
}

func run(srv Server, wg *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			logsagent.Error("Exception has been caught when wait server "+srv.Name(), err)
			wg.Done()
		}
	}()

	go func(srv Server) {
		defer func() {
			if err := recover(); err != nil {
				logsagent.Error("Exception has been caught when run server "+srv.Name(), err)
				wg.Done()
			}
		}()

		srv.Run()
	}(srv)

	<-srv.Stopch()
	wg.Done()
}

//Stop ..
func Stop() {
	for _, srv := range servers {
		go func(srv Server) {
			defer func() {
				if err := recover(); err != nil {
					logsagent.Error("Exception has been caught when stop server " + srv.Name())
					logsagent.Error(err)
				}
			}()

			srv.Stop()
		}(srv)
	}
}
