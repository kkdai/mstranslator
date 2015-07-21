package main

import (
	"fmt"
	"log"
	"os"

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

	//Get Speak audio file stream from server.
	buf, err := msClient.Speak("Returns a wave or mp3 stream of the passed-in text being spoken in the desired language.", "en", "audio/wav")
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("size of buf:", len(buf))
	fo, err := os.Create("speak_output.wav")
	defer fo.Close()
	if _, err := fo.Write(buf); err != nil {
		panic(err)
	}

	//Try to parse input sentence to figure out what language you input.
	retLang, err := msClient.Detect("測試中文")
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("Input Sentence Language:", retLang)

	//Try to parse input sentence to figure out what language you input.
	detectTexts := []string{"哈囉", "あいさつ", "Hello"}
	retLangArray, err := msClient.DetectArray(detectTexts)
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("Input Sentence Language Array:", retLangArray)

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
