package main

import (
	"log"
	"time"

	"github.com/cdvelop/vanify/pkg/sysutil"
)

func main() {
	b := sysutil.BootTime()
	log.Printf("boot: %s", b.Format(time.RFC3339))
	log.Printf("now: %s", time.Now().Format(time.RFC3339))
}
