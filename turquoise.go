package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	ENV_APP_PORT_KEY = "TURQUOISE_APP_PORT"
	APP_NAME         = "turquoise"
	HEALTHY          = "healthy"
)

type Health struct {
	App    string `json:"app"`
	Status string `json:"status"`
}

type Color struct {
	Name string `json:"name,omitempty"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	status := Health{App: APP_NAME, Status: HEALTHY}
	json.NewEncoder(w).Encode(status)
}

func GetColor(w http.ResponseWriter, r *http.Request) {
	version, err := ioutil.ReadFile("version")
	check(err)
	color := Color{Name: string(version)}
	json.NewEncoder(w).Encode(color)
}

func main() {
	port := fmt.Sprintf(":%s", os.Getenv(ENV_APP_PORT_KEY))
	router := mux.NewRouter()
	router.HandleFunc("/turquoise", GetHealth).Methods("GET")
	router.HandleFunc("/turquoise/color", GetColor).Methods("GET")
	log.Fatal(http.ListenAndServe(port, router))
}
