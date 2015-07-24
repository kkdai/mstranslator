package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	ms "github.com/kkdai/mstranslator"
)

func main() {
	//Init new client for mstranslator
	fmt.Println("Connect to MSFT Translator Services..")
	msClient := ms.NewClient("YOUR_CLIEND_ID", "YOUR_CLIENT_SECRET")

	fmt.Println("Input any string for detect and speak:")
	fmt.Printf("Your input string:>")
	scanner := bufio.NewScanner(os.Stdin)
	for !scanner.Scan() {
	}

	inputText := scanner.Text()

	fmt.Println("You input '", inputText, "'")
	retLangCode, err := msClient.Detect(inputText)
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}

	var expectedCodes []string
	expectedCodes = append(expectedCodes, retLangCode)
	retLangName, err := msClient.GetLanguageNames(expectedCodes)
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("We detect it is:", retLangName)

	if retLangCode == "en" {
		retCorrectString, err := msClient.TransformText("en", "general", inputText)
		if err != nil {
			log.Panicf("Error : %s", err.Error())
		}
		fmt.Println("Original: ", inputText, " Correct to:", retCorrectString)
		inputText = retCorrectString
	}

	fmt.Println("We will translate it into ...")

	translationJp, err := msClient.Translate(inputText, retLangCode, "ja")
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("Japanese:>", translationJp)
	fmt.Println("Speaking...")

	buf, err := msClient.Speak(translationJp, "ja", "audio/wav")
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fo, err := os.Create("speak_output_jp.wav")
	defer fo.Close()
	if _, err := fo.Write(buf); err != nil {
		panic(err)
	}
	//speak it out
	err = exec.Command("afplay", "speak_output_jp.wav").Run()

	translationDE, err := msClient.Translate(inputText, retLangCode, "de")
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("Germany:>", translationDE)
	fmt.Println("Speaking...")

	buf2, err := msClient.Speak(translationDE, "de", "audio/wav")
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fo2, err := os.Create("speak_output_de.wav")
	defer fo.Close()
	if _, err := fo2.Write(buf2); err != nil {
		panic(err)
	}
	//speak it out
	err = exec.Command("afplay", "speak_output_de.wav").Run()

	//
	translationEN, err := msClient.Translate(inputText, retLangCode, "en")
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fmt.Println("English:>", translationEN)
	fmt.Println("Speaking...")

	buf3, err := msClient.Speak(translationEN, "en", "audio/wav")
	if err != nil {
		log.Panicf("Error : %s", err.Error())
	}
	fo3, err := os.Create("speak_output_en.wav")
	defer fo.Close()
	if _, err := fo3.Write(buf3); err != nil {
		panic(err)
	}
	//speak it out
	err = exec.Command("afplay", "speak_output_en.wav").Run()
}
