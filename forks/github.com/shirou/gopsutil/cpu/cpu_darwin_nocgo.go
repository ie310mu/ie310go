// +build darwin
// +build !cgo

package cpu

import "github.com/ie310mu/ie310go/forks/github.com/shirou/gopsutil/internal/common"

func perCPUTimes() ([]TimesStat, error) {
	return []TimesStat{}, common.ErrNotImplementedError
}

func allCPUTimes() ([]TimesStat, error) {
	return []TimesStat{}, common.ErrNotImplementedError
}
