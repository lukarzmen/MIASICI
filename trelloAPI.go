package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func pobierzListeZadanZTrello() (listaZadan ListaZadan, err error) {
	trelloAPIResponse, err := http.Get("https://api.trello.com/1/lists/5a51d9e49fedd231a69373fb/cards?oauth_consumer_key=cab0494ae6c25a7525d82da17c07d0b6&oauth_token=e8b62de45abd3f41e6ba4743870cdb8ed8be5c06d9d9aaaa6735bf6aca8b0cac&oauth_signature_method=HMAC-SHA1&oauth_timestamp=1515338724&oauth_nonce=KUo5ur&oauth_version=1.0&oauth_signature=Dh6AJbdi1uQns7ynOBISUmy0EMM=")

	if trelloAPIResponse.StatusCode != http.StatusOK {
		err = errors.New("Kod odpowiedzi rozny od 200")
		return
	}
	defer trelloAPIResponse.Body.Close()
	trelloAPIResponseByte, err := ioutil.ReadAll(trelloAPIResponse.Body)
	if err != nil {
		return
	}

	listaZadan, err = odczytajListeZadan(trelloAPIResponseByte)
	if err != nil {
		return
	}
	return
}

func odczytajListeZadan(trelloApiResponseByte []byte) (listaZadan ListaZadan, err error) {
	var zadania []Zadanie
	err = json.Unmarshal(trelloApiResponseByte, &zadania)
	if err != nil {
		return
	}
	liczbaZadan := len(zadania)

	listaZadan = ListaZadan{
		LiczbaZadan: liczbaZadan,
		Zadania:     zadania,
	}
	return
}
