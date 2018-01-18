package main

type Zadanie struct {
	ID                string `json:"id"`
	Nazwa             string `json:"name"`
	SkroconyURL       string `json:"shortUrl"`
	DataUtworzenia    string `json:"data_utworzenia"`
	TerminZakonczenia string `json:"termin_zakonczenia"`
}

//jak wziąć datę utworzenia zadania
//http://help.trello.com/article/759-getting-the-time-a-card-or-board-was-created
//new Date(1000*parseInt(idBoard.substring(0,8),16));

type ListaZadan struct {
	Zadania                []Zadanie `json:"tasks,omitempty"`
	LiczbaZadan            int       `json:"tasks_count"`
	CzyWystepujaOpoznienia bool      `json:"delay_exists"`
}

type ActivitiTask struct {
	Action    string   `json:"action"`
	Variables []string `json:"variables"`
}

type Build struct {
	Failed    bool
	Branch    string
	StartTime string
	Author    string
}
