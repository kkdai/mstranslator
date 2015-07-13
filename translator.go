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

	"github.com/st3v/tracerr"
)

const (
	API_URL             = "https://datamarket.accesscontrol.windows.net/v2/OAuth2-13/"
	API_SCOPE           = "http://api.microsofttranslator.com"
	ServiceURL          = "http://api.microsofttranslator.com/v2/Http.svc/"
	TranslationURL      = ServiceURL + "Translate"
	DetectURL           = ServiceURL + "Detect"
	GetLanguageNamesURL = ServiceURL + "GetLanguageNames"
)

type Translator struct {
	ClientId     string
	ClientSecret string
	ClientToken  string
}

func doConnect(data url.Values, token string, result interface{}) error {

	client := &http.Client{}
	r, _ := http.NewRequest("POST", API_URL, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Authorization", "auth_token=\""+token+"\"")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)
	return json.Unmarshal(body, &result)
}

func (b *Translator) connect() error {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Add("client_id", b.ClientId)
	data.Add("client_secret", b.ClientSecret)
	data.Add("scope", API_SCOPE)

	result := ResponseToken{}
	doConnect(data, "", &result)
	b.ClientToken = result.AccessToken
	return nil
}

func (b *Translator) Translate(text, from, to string) (string, error) {
	if b.ClientToken == "" {
		b.connect()
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
	request.Header.Add("Authorization", "Bearer "+b.ClientToken)

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

func (b *Translator) Detect(text string) (string, error) {
	if b.ClientToken == "" {
		b.connect()
	}

	uri := fmt.Sprintf(
		"%s?text=%s",
		DetectURL,
		url.QueryEscape(text))

	client := &http.Client{}
	request, err := http.NewRequest("GET", uri, nil)
	request.Header.Add("Content-Type", "text/plain")
	request.Header.Add("Authorization", "Bearer "+b.ClientToken)

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
