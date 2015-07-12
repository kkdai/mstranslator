package main

import ms "../../go-MSTranslator"

func main() {
	msTranslator := ms.Translator{}
	msTranslator.ClientId = ""
	msTranslator.ClientSecret = ""
	msTranslator.Connect()
}
