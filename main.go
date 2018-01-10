package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Serwis Helo Trelo")
	ruter := WezTrasy()
	err := http.ListenAndServe(":80", ruter)
	if err != nil {
		fmt.Errorf(err.Error())
	}

}
