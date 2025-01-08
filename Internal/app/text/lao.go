package text

import (
	"la-skb/pkg"
	"la-skb/pkg/logger"
	"log"
)

var Set map[string]string

func LoadLang() map[string]string {
	lf := "Internal/app/text/locale/lao.lang"
	result, err := pkg.LoadLangFile(lf)
	if err != nil {
		log.Fatalf("Cannot load language file: %s", err)
	}
	return result
}

func InitLang() {
	Set = LoadLang()
	logger.Info("Language file loaded successfully.")
}
