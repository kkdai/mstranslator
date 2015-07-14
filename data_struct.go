package mstranslator

import "encoding/xml"

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
