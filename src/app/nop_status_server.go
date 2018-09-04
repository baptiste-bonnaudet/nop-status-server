package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var STATUS int
var PORT string
var MESSAGE string

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(STATUS)
	fmt.Fprintf(w, MESSAGE)
}

func main() {
	godotenv.Load()
	PORT = os.Getenv("PORT")
	MESSAGE = os.Getenv("MESSAGE")

	var err error

	if STATUS, err = strconv.Atoi(os.Getenv("STATUS")); err != nil {
		log.Fatal("Status is not a number")
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))
}
