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
	chsLock *sync.Mutex
	chs     map[string]chan interface{}
}

//AddChan ..
func (nip *BaseNotiferInProcess) AddChan(key string, ch chan interface{}) {
	nip.chsLock.Lock()
	defer nip.chsLock.Unlock()

	nip.chs[key] = ch
}

//RemoveChan ..
func (nip *BaseNotiferInProcess) RemoveChan(key string, ch chan interface{}) {
	nip.chsLock.Lock()
	defer nip.chsLock.Unlock()

	delete(nip.chs, key)
}

//Name ..
func (nip *BaseNotiferInProcess) Name() string {
	return "BaseNotiferInProcess"
}

//notiferChs ..
func (nip *BaseNotiferInProcess) notiferChs(data interface{}) {
	nip.chsLock.Lock()
	defer nip.chsLock.Unlock()

	for _, ch := range nip.chs {
		ch <- data
	}
}
