package mstranslator

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Authenicator struct {
	ClientId     string
	ClientSecret string
	ClientToken  string
}

func NewAuthenicator(cid, csecret string) *Authenicator {
	return &Authenicator{ClientId: cid, ClientSecret: csecret}
}

func (a *Authenicator) GetToken() string {
	var err error
	if a.ClientToken == "" {
		err = a.authenticate()
	}

	if err != nil {
		return ""
	}

	return a.ClientToken
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

func (a *Authenicator) authenticate() error {
	if a.ClientId == "" || a.ClientSecret == "" {
		return errors.New("No input authenicate information.")
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Add("client_id", a.ClientId)
	data.Add("client_secret", a.ClientSecret)
	data.Add("scope", API_SCOPE)

	result := ResponseToken{}
	doAuthenticate(data, "", &result)
	a.ClientToken = result.AccessToken
	return nil
}
