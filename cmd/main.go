package main

import (
	"log"
	"net/http"
	"task/internal/app"
)

func main() {
	router := app.Run()

	log.Println("Server is running http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Println("Server is busy...")
		return
	}
}
