package main

type Zadanie struct {
	ID          string `json:"id"`
	Nazwa       string `json:"name"`
	SkroconyURL string `json:"shortUrl"`
}

type ListaZadan struct {
	Zadania     []Zadanie `json:"tasks,omitempty"`
	LiczbaZadan int       `json:"tasks_count"`
}
