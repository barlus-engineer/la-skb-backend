package text

import (
	"la-skb/pkg"
	"la-skb/pkg/logger"
	"log"
)

var Set map[string]string

func LoadLang() map[string]string {
	lf := "locale/lao.lang"
	result, err := pkg.LoadLangFile(lf)
	if err != nil {
		log.Fatalf("Cannot load text file: %s", err)
	}
	return result
}

func InitLang() {
	Set = LoadLang()
	logger.Info("Text file loaded successfully.")
}
