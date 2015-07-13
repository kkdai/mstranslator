package main

import (
	"fmt"
	"log"

	ms "../../go-MSTranslator"
)

func main() {
	msTranslator := ms.Translator{}
	msTranslator.ClientId = ""
	msTranslator.ClientSecret = ""

	translation, err := msTranslator.Translate("Hello World!", "en", "de")
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println(translation)

	retLang, err := msTranslator.Detect("測試中文")
	fmt.Println("Lang=", retLang)
}
