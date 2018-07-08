package logs

import (
	"log"
	"strings"
)

func PrintTitle(key string, v ...string) {
	log.Printf("%s [%s] ==========================", key, strings.Join(v, ","))
}
