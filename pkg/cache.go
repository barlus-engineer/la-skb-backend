package pkg

import (
	"strings"
)

func RDBstringify(d ...string) string {
	st := strings.Join(d, " ")
	return st
}

func RDBpaser(d string) []string {
	result := strings.Fields(d)
	return result
}