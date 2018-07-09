package domain

import (
	"log"
)

type TaskResult struct {
	Target string
	Rlt    string
	Msg    interface{}
}

func (this *TaskResult) Out() {
	log.Printf("%s|%s => \n{\nmsg:%s}\n", this.Target, this.Rlt, this.Msg)
}
