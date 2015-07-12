package mstranslator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const (
	API_URL   = "https://datamarket.accesscontrol.windows.net/v2/OAuth2-13/"
	API_SCOPE = "http://api.microsofttranslator.com"
)

type Translator struct {
	ClientId     string
	ClientSecret string
	ClientToken  string
}

func DoConnect(data url.Values, token string, result interface{}) error {

	client := &http.Client{}
	r, _ := http.NewRequest("POST", API_URL, bytes.NewBufferString(data.Encode())) // <-- URL-encoded payload
	r.Header.Add("Authorization", "auth_token=\""+token+"\"")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))

	return json.Unmarshal(body, &result)
}

func (b *Translator) Connect() error {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Add("client_id", b.ClientId)
	data.Add("client_secret", b.ClientSecret)
	data.Add("scope", API_SCOPE)

	result := ResponseToken{}
	DoConnect(data, "", &result)
	b.ClientToken = result.AccessToken
	fmt.Println("ret:", result)
	return nil
}

func (b *Translator) Translate(text, from, to string) (string, error) {
	return "", nil
}
