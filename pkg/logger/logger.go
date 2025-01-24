package logger

import (
	"fmt"
	"la-skb/pkg/color"
	"log"
)


func Info(value string) {
	color := colors.Blue
	text := fmt.Sprintf("[%sInfo\033[0m] %s", color, value)
	log.Println(text)
}

func Warning(value string) {
	color := colors.Yellow
	text := fmt.Sprintf("[%sWarning\033[0m] %s", color, value)
	log.Println(text)
}

func Alert(value string) {
	color := colors.Red
	text := fmt.Sprintf("%sAlert\033[0m] %s", color, value)
	log.Println(text)
}