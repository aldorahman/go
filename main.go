package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("listening on port 3000...")
	http.ListenAndServe(":"+port, nil)
}
