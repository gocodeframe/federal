package domain

import (
	"log"
)

type TaskResult struct {
	Target string
	Rlt    string
	Msg    string
}

func (this *TaskResult) Out() {
	log.Printf("%s | %s ==> {msg:%s}\n", this.Target, this.Rlt, this.Msg)
}
