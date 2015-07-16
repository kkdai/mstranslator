package mstranslator

//Store all API URI address.
const (
	API_URL   = "https://datamarket.accesscontrol.windows.net/v2/OAuth2-13/"
	API_SCOPE = "http://api.microsofttranslator.com"

	TransformTextURL = "http://api.microsofttranslator.com/V3/json/TransformText"

	ServiceURL                  = "http://api.microsofttranslator.com/v2/Http.svc/"
	TranslationURL              = ServiceURL + "Translate"
	GetTranslationsURL          = ServiceURL + "GetTranslations"
	DetectURL                   = ServiceURL + "Detect"
	DetectArrayURL              = ServiceURL + "DetectArray"
	SpeakURL                    = ServiceURL + "Speak"
	GetLanguageNamesURL         = ServiceURL + "GetLanguageNames"
	GetLanguagesForTranslateURL = ServiceURL + "GetLanguagesForTranslate"
	GetLanguagesForSpeakURL     = ServiceURL + "GetLanguagesForSpeak"
)
