package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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
	go zatwierdzOdczytanieListyZdan(23)
	return
}

func odczytajListeZadan(trelloApiResponseByte []byte) (listaZadan ListaZadan, err error) {
	var zadania []Zadanie
	err = json.Unmarshal(trelloApiResponseByte, &zadania)
	if err != nil {
		return
	}

	listaZadan = przetworzZadaniaDoListy(zadania)
	return
}

func przetworzZadaniaDoListy(zadania []Zadanie) (listaZadan ListaZadan) {
	var zadaniaZData []Zadanie
	czyOpoznienie := false
	for _, zadanie := range zadania {
		zadanieZData, czyOpoznienieZadania := obliczDateUtworzeniaUlotki(zadanie)
		zadaniaZData = append(zadaniaZData, zadanieZData)
		if czyOpoznienieZadania {
			czyOpoznienie = true
		}
	}

	liczbaZadan := len(zadania)

	listaZadan = ListaZadan{
		LiczbaZadan:            liczbaZadan,
		Zadania:                zadaniaZData,
		CzyWystepujaOpoznienia: czyOpoznienie,
	}
	return
}
func obliczDateUtworzeniaUlotki(zadanie Zadanie) (zadaniaZData Zadanie, czyOpoznienie bool) {
	id := zadanie.ID
	hexTimestamp := id[0:8]

	sec, err := strconv.ParseInt(hexTimestamp, 16, 64)
	if err != nil {
		panic(err)
	}
	creationDateTimeStamp := time.Unix(sec, 0)
	creationDate := creationDateTimeStamp.Format(time.RFC3339)
	dedlineDate := creationDateTimeStamp.AddDate(0, 0, 2)
	czyOpoznienie = dedlineDate.Before(time.Now())

	terminZakonczeniaString := dedlineDate.Format(time.RFC3339)
	zadanie.DataUtworzenia = creationDate
	zadanie.TerminZakonczenia = terminZakonczeniaString
	zadaniaZData = zadanie
	return
}

func zatwierdzOdczytanieListyZdan(taskID int) (err error) {
	completeTaskAction := ActivitiTask{
		Action:    "complete",
		Variables: []string{},
	}
	completeTaskActionJSON, err := json.Marshal(&completeTaskAction)

	client := &http.Client{}

	/* Authenticate */
	taksIDString := strconv.Itoa(taskID)
	request, err := http.NewRequest(http.MethodPost, "http://80.211.255.185:32771/activiti-rest/service/runtime/tasks/"+taksIDString, bytes.NewBuffer(completeTaskActionJSON))
	request.SetBasicAuth("kermit", "kermit")
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	defer response.Body.Close()
	if response.StatusCode != 200 {
		err = errors.New("Nie można zatwierdzić zadania odczytania listy zadan w Activiti")
	}

	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
