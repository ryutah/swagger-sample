package main

import (
	"log"
	"net/http"

	sw "github.com/ryutah/swagger-sample/swagger"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
