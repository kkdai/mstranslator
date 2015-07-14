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
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("Lang=", retLang)

	retSupportLangcodes, err := msTranslator.GetLanguagesForTranslate()
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("LangCodes=", retSupportLangcodes)

	expectedCodes := []string{"en", "de", "es", "ru", "jp"}
	retLangName, err := msTranslator.GetLanguageNames(expectedCodes)
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("Lang=", retLangName)
}
