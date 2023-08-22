package main

import (
	"log"
	"net/http"
	"task/internal/app"
)

func main() {
	router := app.Run()

	log.Println("Server is running http://localhost:7777")
	if err := http.ListenAndServe(":7777", router); err != nil {
		log.Println("Server is busy...")
		return
	}
}
