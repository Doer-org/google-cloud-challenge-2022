package logger

import (
	"log"
	"time"
)

func Println(message string) {
	log.Printf(
		"[server log, %s] : %s",
		time.Now().String(),
		message,
	)
}
