package ie310go

import (
	"github.com/ie310mu/ie310go/session"
)

//SetDefaultSessionInfo ..
func SetDefaultSessionInfo(scn string, esc bool) {
	sessionCookieName = scn
	enableSetCookie = esc
}

//AppSessionInitFunc session初始化***********不指定时使用默认的内存管理方式
var AppSessionInitFunc appInitFunc = initAppSessionAsDefault

//session初始化
func initAppSessionAsDefault() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      sessionCookieName,
		EnableSetCookie: enableSetCookie,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	session.GlobalSessions, _ = session.NewManager("memory", sessionConfig)
	go session.GlobalSessions.GC()
}

var sessionCookieName = "ie310gosessionid"
var enableSetCookie = true
