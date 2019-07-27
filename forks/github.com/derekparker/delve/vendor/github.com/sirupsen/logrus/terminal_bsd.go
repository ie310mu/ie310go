// +build darwin freebsd openbsd netbsd dragonfly
// +build !appengine,!gopherjs

package logrus

import "github.com/ie310mu/ie310go/forks/golang.org/x/sys/unix"

const ioctlReadTermios = unix.TIOCGETA

type Termios unix.Termios
