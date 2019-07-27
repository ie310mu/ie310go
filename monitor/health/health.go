package health

import (
	"fmt"
)

var nameChecks = []*nameCheck{}

//RegisteHealthCheck ..
func RegisteHealthCheck(n string, s IHealthCheck) {
	nc := &nameCheck{name: n, check: s}
	nameChecks = append(nameChecks, nc)
}

//IHealthCheck ..
type IHealthCheck interface {
	Check() error
}

type nameCheck struct {
	name  string
	check IHealthCheck
}

//Result ..
type Result struct {
	Name   string `json:"name"    `
	Result string `json:"result"    `
}

//Check ..
func Check() []*Result {
	rs := []*Result{}
	for _, nameCheck := range nameChecks {
		err := nameCheck.check.Check()
		hr := &Result{Name: nameCheck.name, Result: "success"}
		if err != nil {
			hr.Result = fmt.Sprintf("%v", err)
		}
		rs = append(rs, hr)
	}

	return rs
}
