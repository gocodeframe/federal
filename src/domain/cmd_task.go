package domain

import (
	"infra/logs"
)

type Task struct {
	Name string
	Cmds []string
	rlt  TaskResult
}

func (this *Task) Exec(task string, host string) error {
	logs.PrintTitle("Task", host)
	this.rlt.Out()
	return nil
}
