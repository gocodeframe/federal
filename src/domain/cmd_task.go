package domain

import (
	"infra/logs"
)

type Task struct {
	Name string
	Cmds []string
}

func (this *Task) Exec(task string, host string) error {
	logs.PrintTitle("Task", host)
	return nil
}
