package logger

import (
	"fmt"
	"la-skb/pkg"
	"log"
)


func Info(value string) {
	color := pkg.Colors.Blue
	text := fmt.Sprintf("[%sInfo\033[0m] %s", color, value)
	log.Println(text)
}

func Warning(value string) {
	color := pkg.Colors.Yellow
	text := fmt.Sprintf("[%sWarning\033[0m] %s", color, value)
	log.Println(text)
}

func Alert(value string) {
	color := pkg.Colors.Red
	text := fmt.Sprintf("%sAlert\033[0m] %s", color, value)
	log.Println(text)
}