package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func pobierzInformacjeOBuildachZCircleCI() (listaBuildow ListaBuildow, err error) {
	circleCIResponse, err := http.Get("https://circleci.com/api/v1.1/project/github/lukarzmen/miasi-serwis?circle-token=b8f2bc36d1fbefb2bc1f043c157c49cc0c334fb2")

	if circleCIResponse.StatusCode != http.StatusOK {
		err = errors.New("Kod odpowiedzi CircleCIAPI  rozny od 200")
		return
	}
	defer circleCIResponse.Body.Close()
	circleCIResponseByte, err := ioutil.ReadAll(circleCIResponse.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(circleCIResponseByte, &listaBuildow)
	if err != nil {
		return
	}

	go zatwierdzOdczytanieListyZdan(24)
	return
}
