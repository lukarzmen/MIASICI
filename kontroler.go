package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func pobierzListeZadan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/pobierzListeZadan")

	listaZadan, err := pobierzListeZadanZTrello()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if listaZadan.CzyWystepujaOpoznienia {
		fmt.Println("Wystepuja opoznienia!")
	}
	fmt.Println("Liczba zadan: ", listaZadan.LiczbaZadan)

	listaZadanByte, err := json.Marshal(listaZadan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(listaZadanByte)
}

func pobierzBuildy(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/pobierzBuildy")

	listaBuildow, err := pobierzInformacjeOBuildachZCircleCI()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(listaBuildow.ListaBuildow) < 1 {
		err = errors.New("Brak budowanych zadan")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if listaBuildow.ListaBuildow[0].Failed == nil {
		err = errors.New("Zbudowanie paczki nie powiodlo sie")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Liczba buildow: ", len(listaBuildow.ListaBuildow))

	listaBuildowByte, err := json.Marshal(listaBuildow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(listaBuildowByte)
}
