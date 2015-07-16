package mstranslator

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/st3v/tracerr"
)

type LanguageProvider struct {
	authenicator *Authenicator
}

func NewLanguageProvider(auth *Authenicator) *LanguageProvider {
	return &LanguageProvider{authenicator: auth}
}

func getXMLArrayFromString(values []string) *ResponseArray {
	return &ResponseArray{
		Namespace:         "http://schemas.microsoft.com/2003/10/Serialization/Arrays",
		InstanceNamespace: "http://www.w3.org/2001/XMLSchema-instance",
		Strings:           values,
	}
}

func (l *LanguageProvider) Detect(text string) (string, error) {
	token := l.authenicator.GetToken()

	uri := fmt.Sprintf(
		"%s?text=%s",
		DetectURL,
		url.QueryEscape(text))

	client := &http.Client{}
	request, err := http.NewRequest("GET", uri, nil)
	request.Header.Add("Content-Type", "text/plain")
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := client.Do(request)
	if err != nil {
		return "", tracerr.Wrap(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", tracerr.Wrap(err)
	}

	retDetect := &ResponseXML{}
	err = xml.Unmarshal(body, &retDetect)
	if err != nil {
		return "", tracerr.Wrap(err)
	}

	return retDetect.Value, nil
}

func (l *LanguageProvider) DetectArray(text []string) ([]string, error) {
	token := l.authenicator.GetToken()

	payload, _ := xml.Marshal(getXMLArrayFromString(text))

	client := &http.Client{}
	request, err := http.NewRequest("POST", DetectArrayURL, strings.NewReader(string(payload)))
	request.Header.Add("Content-Type", "text/xml")
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := client.Do(request)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	retDetect := &ResponseArray{}
	err = xml.Unmarshal(body, &retDetect)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return retDetect.Strings, nil
}

func (l *LanguageProvider) GetTranslations(text, from, to string, maxTranslations int) ([]ResponseTranslationMatch, error) {
	token := l.authenicator.GetToken()

	uri := fmt.Sprintf(
		"%s?text=%s&from=%s&to=%s&maxTranslations=%d",
		GetTranslationsURL,
		url.QueryEscape(text),
		url.QueryEscape(from),
		url.QueryEscape(to),
		maxTranslations)

	client := &http.Client{}
	request, err := http.NewRequest("POST", uri, nil)
	request.Header.Add("Content-Type", "text/xml")
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := client.Do(request)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	xmlResult := GetTranslationsResponse{}
	err = xml.Unmarshal(body, &xmlResult)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return xmlResult.Translations.TransMatch, nil
}

//Retrieves friendly names for the languages passed in as the parameter languageCodes, and localized using the passed locale language.
func (l *LanguageProvider) GetLanguageNames(codes []string) ([]string, error) {
	token := l.authenicator.GetToken()

	payload, _ := xml.Marshal(getXMLArrayFromString(codes))
	uri := fmt.Sprintf("%s?locale=en", GetLanguageNamesURL)

	client := &http.Client{}
	request, err := http.NewRequest("POST", uri, strings.NewReader(string(payload)))
	request.Header.Add("Content-Type", "text/xml")
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := client.Do(request)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	retLangs := &ResponseArray{}
	err = xml.Unmarshal(body, &retLangs)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return retLangs.Strings, nil
}

func (l *LanguageProvider) GetLanguagesForTranslate() ([]string, error) {
	token := l.authenicator.GetToken()

	client := &http.Client{}
	request, err := http.NewRequest("GET", GetLanguagesForTranslateURL, nil)
	request.Header.Add("Content-Type", "text/plain")
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := client.Do(request)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	retLangs := &ResponseArray{}
	err = xml.Unmarshal(body, &retLangs)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return retLangs.Strings, nil
}

func (l *LanguageProvider) GetLanguagesForSpeak() ([]string, error) {
	token := l.authenicator.GetToken()

	client := &http.Client{}
	request, err := http.NewRequest("GET", GetLanguagesForSpeakURL, nil)
	request.Header.Add("Content-Type", "text/plain")
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := client.Do(request)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	retLangs := &ResponseArray{}
	err = xml.Unmarshal(body, &retLangs)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return retLangs.Strings, nil
}
