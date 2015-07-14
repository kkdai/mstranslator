package mstranslator

import "encoding/xml"

type XMLIntValue struct {
	Text int `xml:",chardata"`
}

type XMLStringValue struct {
	Text string `xml:",chardata"`
}

type GetTranslationsResponse struct {
	XMLName      xml.Name             `xml:"GetTranslationsResponse"`
	Translations ResponseTranslations `xml:"Translations"`
}

type ResponseTranslations struct {
	TransMatch []ResponseTranslationMatch `xml:"TranslationMatch"`
}

type ResponseTranslationMatch struct {
	Count               XMLIntValue    `xml:"Count"`
	MatchDegree         XMLIntValue    `xml:"MatchDegree"`
	MatchedOriginalText XMLStringValue `xml:"MatchedOriginalText"`
	Rating              XMLIntValue    `xml:"Rating"`
	TranslatedText      XMLStringValue `xml:"TranslatedText"`
}

type ResponseToken struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	Scope       string `json:"scope"`
}

type TransformTextResponse struct {
	ErrorCondition   int    `json:"ec"`       // A positive number representing an error condition
	ErrorDescriptive string `json:"em"`       // A descriptive error message
	Sentence         string `json:"sentence"` // transformed text
}

type ResponseXML struct {
	XMLName   xml.Name `xml:"string"`
	Namespace string   `xml:"xmlns,attr"`
	Value     string   `xml:",innerxml"`
}

type ResponseArray struct {
	XMLName           xml.Name `xml:"ArrayOfstring"`
	Namespace         string   `xml:"xmlns,attr"`
	InstanceNamespace string   `xml:"xmlns:i,attr"`
	Strings           []string `xml:"string"`
}
