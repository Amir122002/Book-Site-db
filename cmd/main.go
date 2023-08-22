package main

import (
	"book/internal/configs"
	"book/internal/handlers"
	"log"
	"net/http"
)

func main() {
	config, err := configs.InitConfig()
	if err != nil {
		log.Fatal()
	}

	address := config.Host + config.Port
	router := handlers.InitRouter()

	srv := http.Server{
		Addr:    address,
		Handler: router,
	}
	log.Print("start")
	err = srv.ListenAndServe()
	if err != nil {
		log.Println("listening and service error:", err)
	}
}
