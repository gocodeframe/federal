package domain

import (
	"infra/logs"
)

type Task struct {
	Name string `yaml:"name"`
	Cmd  string `yaml:"cmd"`
	rlt  TaskResult
}

func (this *Task) Exec(task string, host string) error {
	logs.PrintTitle("Task", host)
	(&TaskResult{Rlt: "success", Target: "all", Msg: "hello world"}).Out()
	return nil
}
