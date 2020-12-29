package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	var port = ":8889"
	if os.Getenv("PORT") != "" {
		port = fmt.Sprintf(":%v", os.Getenv("PORT"))
	}
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	log.Printf("go-web-hostname running on port%v", port)
	log.Fatal(http.ListenAndServe(port, r))
}

// HomeHandler ..
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	res, err := getHostname()
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(res)
}

func getHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return hostname, nil
}
