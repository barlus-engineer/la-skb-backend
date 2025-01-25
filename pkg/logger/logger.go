package logger

import (
	"fmt"
	"la-skb/pkg/color"
	"log"
)


func Info(values ...any) {
	value := fmt.Sprintf("%v", values)
	value = value[1:len(value)-1]
	color := colors.Blue
	text := fmt.Sprintf("[%sInfo\033[0m]\t%s", color, value)
	log.Println(text)
}

func Warning(values ...any) {
	value := fmt.Sprintf("%v", values)
	value = value[1:len(value)-1]
	color := colors.Yellow
	text := fmt.Sprintf("[%sWarning\033[0m]\t%s", color, value)
	log.Println(text)
}

func Alert(values ...any) {
	value := fmt.Sprintf("%v", values)
	value = value[1:len(value)-1]
	color := colors.Red
	text := fmt.Sprintf("[%sAlert\033[0m]\t%s", color, value)
	log.Println(text)
}