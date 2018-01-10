package main

type Zadanie struct {
	ID          string `json:"id"`
	Nazwa       string `json:"name"`
	SkroconyURL string `json:"shortUrl"`
}

//jak wziąć datę utworzenia zadania
//http://help.trello.com/article/759-getting-the-time-a-card-or-board-was-created
//new Date(1000*parseInt(idBoard.substring(0,8),16));

type ListaZadan struct {
	Zadania     []Zadanie `json:"tasks,omitempty"`
	LiczbaZadan int       `json:"tasks_count"`
}
