package debugger

import (
	"fmt"
	sys "github.com/ie310mu/ie310go/forks/golang.org/x/sys/unix"
)

func attachErrorMessage(pid int, err error) error {
	//TODO: mention certificates?
	return fmt.Errorf("could not attach to pid %d: %s", pid, err)
}

func stopProcess(pid int) error {
	return sys.Kill(pid, sys.SIGSTOP)
}
