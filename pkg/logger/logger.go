package logger

import (
	"fmt"
	"log"
)

func Info(value string) {
	text := fmt.Sprintf("[\033[34mInfo\033[0m]	%s", value)
	log.Println(text)
}

func Warning(value string) {
	text := fmt.Sprintf("[\033[33mWarning\033[0m]	%s", value)
	log.Println(text)
}

func Alert(value string) {
	text := fmt.Sprintf("[\033[31mAlert\033[0m]	%s", value)
	log.Println(text)
}