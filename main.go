package main

import (
	"api"
	"log"
)

func main() {
	log.Printf("===========ferderal ============")
	(&api.CmdApi{}).RunCmd()
}
