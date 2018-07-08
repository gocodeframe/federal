package domain

import (
	"infra/logs"
)

type ITask interface {
	Exec(task string, host string) error
}

type Assigment struct {
	Hosts    []string
	TaskList []ITask
}

func (this *Assigment) Run() {
	logs.PrintTitle("ASSIGN", this.Hosts...)
	for _, task := range this.TaskList {
		for _, host := range this.Hosts {
			task.Exec("", host)
		}
	}
}
