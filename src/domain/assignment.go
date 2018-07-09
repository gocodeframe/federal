package domain

import (
	"infra/logs"
)

type Assigment struct {
	Hosts    []string
	TaskList []Task
}

func (this *Assigment) Run() {
	logs.PrintTitle("ASSIGN", this.Hosts...)
	for _, task := range this.TaskList {
		for _, host := range this.Hosts {
			task.Exec("", host)
		}
	}
}
