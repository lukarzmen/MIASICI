package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func WezTrasy() (r *mux.Router) {
	r = mux.NewRouter()
	for _, trasa := range trasy {
		r.HandleFunc(trasa.Sciezka, trasa.Akcja)
	}
	return
}

var trasy []Trasa = []Trasa{
	Trasa{
		Sciezka: "/lista-zadan",
		Akcja:   pobierzListeZadan,
	},
	Trasa{
		Sciezka: "/lista-buildow",
		Akcja:   pobierzBuildy,
	},
}

type Trasa struct {
	Sciezka string
	Akcja   func(w http.ResponseWriter, r *http.Request)
}
