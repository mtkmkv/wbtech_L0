package app

import (
	"L0/internal/database"
	"L0/internal/services"
	"L0/internal/views"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	err := database.SyncCacheAndDatabase(database.Connection())
	if err != nil {
		log.Print(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", views.MainPage)
	router.HandleFunc("/message/{uid:[0-9, A-Z, a-z]+}", views.MessageHandler)

	go services.GetMessage()

	err = http.ListenAndServe(views.Port, router)
	if err != nil {
		log.Printf("Error in lauching server: %v", err)
	}
}
