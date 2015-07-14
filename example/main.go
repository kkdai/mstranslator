package main

import (
	"fmt"
	"log"

	ms "github.com/kkdai/mstranslator"
)

func main() {
	msTranslator := ms.Translator{}
	msTranslator.ClientId = ""
	msTranslator.ClientSecret = ""

	//Translate "Hello World" from English to France.
	translation, err := msTranslator.Translate("Hello World!", "en", "de")
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println(translation)

	//Try to parse input sentence to figure out what language you input.
	retLang, err := msTranslator.Detect("測試中文")
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("Inpurt Sentence Language:", retLang)

	//Get all language support by Microsoft Translator.
	retSupportLangcodes, err := msTranslator.GetLanguagesForTranslate()
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("Server Support Language Codes:", retSupportLangcodes)

	//Get detail Language Name (ex: en -> English)
	expectedCodes := []string{"en", "de", "es", "ru", "jp"}
	retLangName, err := msTranslator.GetLanguageNames(expectedCodes)
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("Inpurt Full Language Name:", retLangName)

	//Correct senstence should be "This is too strange i just want to go home soon".
	oriSentence := "Dis is 2 strange i juss wanna go home sooooooon"
	retCorrectString, err := msTranslator.TransformText("en", "general", oriSentence)
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("Original: ", oriSentence, " Correct to:", retCorrectString)

	//Find a possible translation result for whole setence
	retGet, err := msTranslator.GetTranslations("una importante contribución a la rentabilidad de la empresa", "es", "en", 5)
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println(retGet)

}
