package route

import "sync"

type NotifierInProcess interface {
	Name() string
	AddChan(key string, ch chan interface{})
	RemoveChan(key string, ch chan interface{})
	TryRun()
	Stop()
}

type BaseNotiferInProcess struct {
	ChsLock *sync.Mutex
	Chs     map[string]chan interface{}
}

//AddChan ..
func (nip *BaseNotiferInProcess) AddChan(key string, ch chan interface{}) {
	nip.ChsLock.Lock()
	defer nip.ChsLock.Unlock()

	nip.Chs[key] = ch
}

//RemoveChan ..
func (nip *BaseNotiferInProcess) RemoveChan(key string, ch chan interface{}) {
	nip.ChsLock.Lock()
	defer nip.ChsLock.Unlock()

	delete(nip.Chs, key)
}

//Name ..
func (nip *BaseNotiferInProcess) Name() string {
	return "BaseNotiferInProcess"
}

//notiferChs ..
func (nip *BaseNotiferInProcess) NotiferChs(data interface{}) {
	nip.ChsLock.Lock()
	defer nip.ChsLock.Unlock()

	for _, ch := range nip.Chs {
		ch <- data
	}
}
