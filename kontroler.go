package main

import (
	"encoding/json"
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
