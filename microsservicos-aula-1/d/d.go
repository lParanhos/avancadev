package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Status string
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	response := Response{Status: "Great, finish here :)"}

	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal("Error on json parse")
	}

	fmt.Fprintf(w, string(jsonData))

}
