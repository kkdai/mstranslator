package mstranslator

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/st3v/tracerr"
)

const (
	API_URL   = "https://datamarket.accesscontrol.windows.net/v2/OAuth2-13/"
	API_SCOPE = "http://api.microsofttranslator.com"

	TransformTextURL = "http://api.microsofttranslator.com/V3/json/TransformText"

	ServiceURL                  = "http://api.microsofttranslator.com/v2/Http.svc/"
	TranslationURL              = ServiceURL + "Translate"
	GetTranslationsURL          = ServiceURL + "GetTranslations"
	TranslateArray              = ServiceURL + "TranslateArray"
	DetectURL                   = ServiceURL + "Detect"
	GetLanguageNamesURL         = ServiceURL + "GetLanguageNames"
	GetLanguagesForTranslateURL = ServiceURL + "GetLanguagesForTranslate"
)

type Translator struct {
	ClientId     string
	ClientSecret string
	ClientToken  string
}

func getXMLArrayFromString(values []string) *ResponseArray {
	return &ResponseArray{
		Namespace:         "http://schemas.microsoft.com/2003/10/Serialization/Arrays",
		InstanceNamespace: "http://www.w3.org/2001/XMLSchema-instance",
		Strings:           values,
	}
}

func doAuthenticate(data url.Values, token string, result interface{}) error {

	client := &http.Client{}
	r, _ := http.NewRequest("POST", API_URL, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Authorization", "auth_token=\""+token+"\"")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)
	return json.Unmarshal(body, &result)
}

func (t *Translator) authenticate() error {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Add("client_id", t.ClientId)
	data.Add("client_secret", t.ClientSecret)
	data.Add("scope", API_SCOPE)

	result := ResponseToken{}
	doAuthenticate(data, "", &result)
	t.ClientToken = result.AccessToken
	return nil
}

//Translates a text string from one language to another.
func (t *Translator) Translate(text, from, to string) (string, error) {
	if t.ClientToken == "" {
		t.authenticate()
	}

	uri := fmt.Sprintf(
		"%s?text=%s&from=%s&to=%s",
		TranslationURL,
		url.QueryEscape(text),
		url.QueryEscape(from),
		url.QueryEscape(to))

	client := &http.Client{}
	request, err := http.NewRequest("GET", uri, nil)
	request.Header.Add("Content-Type", "text/plain")
	request.Header.Add("Authorization", "Bearer "+t.ClientToken)

	response, err := client.Do(request)
	if err != nil {
		return "", tracerr.Wrap(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", tracerr.Wrap(err)
	}

	translation := &ResponseXML{}
	err = xml.Unmarshal(body, &translation)
	if err != nil {
		return "", tracerr.Wrap(err)
	}

	return translation.Value, nil
}

//The TransformText method is a text normalization function for social media, which returns a normalized form of the input.
//The method can be used as a preprocessing step in Machine Translation or other applications, which expect clean input text than is typically found in social media or user-generated content. The function currently works only with English input.
func (t *Translator) TransformText(lang, category, text string) (string, error) {
	if t.ClientToken == "" {
		t.authenticate()
	}

	uri := fmt.Sprintf(
		"%s?sentence=%s&category=%s&language=%s",
		TransformTextURL,
		url.QueryEscape(text),
		url.QueryEscape(category),
		url.QueryEscape(lang))

	client := &http.Client{}
	request, err := http.NewRequest("GET", uri, nil)
	request.Header.Add("Content-Type", "text/plain")
	request.Header.Add("Authorization", "Bearer "+t.ClientToken)

	response, err := client.Do(request)
	if err != nil {
		return "", tracerr.Wrap(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", tracerr.Wrap(err)
	}

	// Microsoft Server json response contain BOM, need to trim.
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	transTransform := TransformTextResponse{}
	err = json.Unmarshal(body, &transTransform)
	if err != nil {
		return "", tracerr.Wrap(err)
	}

	return transTransform.Sentence, nil
}

// Use the Detect Method to identify the language of a selected piece of text.
func (t *Translator) Detect(text string) (string, error) {
	if t.ClientToken == "" {
		t.authenticate()
	}

	uri := fmt.Sprintf(
		"%s?text=%s",
		DetectURL,
		url.QueryEscape(text))

	client := &http.Client{}
	request, err := http.NewRequest("GET", uri, nil)
	request.Header.Add("Content-Type", "text/plain")
	request.Header.Add("Authorization", "Bearer "+t.ClientToken)

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

//Retrieves an array of translations for a given language pair from the store and the MT engine. GetTranslations differs from Translate as it returns all available translations.
func (t *Translator) GetTranslations(text, from, to string, maxTranslations int) ([]ResponseTranslationMatch, error) {
	if t.ClientToken == "" {
		t.authenticate()
	}

	// payload, _ := xml.Marshal(getXMLArrayFromString(codes))
	uri := fmt.Sprintf(
		"%s?text=%s&from=%s&to=%s&maxTranslations=%d",
		GetTranslationsURL,
		url.QueryEscape(text),
		url.QueryEscape(from),
		url.QueryEscape(to),
		maxTranslations)

	client := &http.Client{}
	request, err := http.NewRequest("POST", uri, nil) // strings.NewReader(string(payload)))
	request.Header.Add("Content-Type", "text/xml")
	request.Header.Add("Authorization", "Bearer "+t.ClientToken)

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
func (t *Translator) GetLanguageNames(codes []string) ([]string, error) {
	if t.ClientToken == "" {
		t.authenticate()
	}

	payload, _ := xml.Marshal(getXMLArrayFromString(codes))
	uri := fmt.Sprintf("%s?locale=en", GetLanguageNamesURL)

	client := &http.Client{}
	request, err := http.NewRequest("POST", uri, strings.NewReader(string(payload)))
	request.Header.Add("Content-Type", "text/xml")
	request.Header.Add("Authorization", "Bearer "+t.ClientToken)

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

func (t *Translator) GetLanguagesForTranslate() ([]string, error) {
	if t.ClientToken == "" {
		t.authenticate()
	}

	client := &http.Client{}
	request, err := http.NewRequest("GET", GetLanguagesForTranslateURL, nil)
	request.Header.Add("Content-Type", "text/plain")
	request.Header.Add("Authorization", "Bearer "+t.ClientToken)

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
