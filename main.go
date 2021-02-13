package main

import (
	"log"
	"net/http"

	"github.com/ahmedshaaban/home24/handlers"
	"github.com/ahmedshaaban/home24/services"
)

func main() {
	sService := services.NewScrapperService()
	sHandler := handlers.NewScrapperHandler(sService, http.DefaultClient)
	http.HandleFunc("/scrap", sHandler.Scrap)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
