package main

import (
	"fmt"
	"log"

	ms "github.com/kkdai/mstranslator"
)

func main() {
	//Init new client for mstranslator
	msClient := ms.NewClient("", "")

	//Translate "Hello World" from English to France.
	translation, err := msClient.Translate("Hello World!", "en", "de")
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println(translation)

	//Try to parse input sentence to figure out what language you input.
	retLang, err := msClient.Detect("測試中文")
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("Inpurt Sentence Language:", retLang)

	//Get all language support by Microsoft Translator.
	retSupportLangcodes, err := msClient.GetLanguagesForTranslate()
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("Server Support Language Codes:", retSupportLangcodes)

	//Get all language support speech by Microsoft Translator.
	retSupportSpeakcodes, err := msClient.GetLanguagesForSpeak()
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("Server Support Speak Language Codes:", retSupportSpeakcodes)

	//Get detail Language Name (ex: en -> English)
	expectedCodes := []string{"en", "de", "es", "ru", "jp"}
	retLangName, err := msClient.GetLanguageNames(expectedCodes)
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("Inpurt Full Language Name:", retLangName)

	//Correct senstence should be "This is too strange i just want to go home soon".
	oriSentence := "Dis is 2 strange i juss wanna go home sooooooon"
	retCorrectString, err := msClient.TransformText("en", "general", oriSentence)
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("Original: ", oriSentence, " Correct to:", retCorrectString)

	//Find a possible translation result for whole setence
	retGet, err := msClient.GetTranslations("una importante contribución a la rentabilidad de la empresa", "es", "en", 5)
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println(retGet)

}
