package pc

import (
	"os/exec"
	"runtime"

	"github.com/ie310mu/ie310go/common/throw"
)

var commands = map[string]string{
	"darwin": "open",
	"linux":  "xdg-open",
}

func OpenUrl(uri string) {
	os := runtime.GOOS
	if os == "windows" {
		err := exec.Command("cmd", "/c", "start", uri).Start()
		throw.CheckErr(err)
	} else {
		run, _ := commands[runtime.GOOS]
		err := exec.Command(run, uri).Start()
		throw.CheckErr(err)
	}
}
