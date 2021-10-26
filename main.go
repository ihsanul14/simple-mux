package main

import (
	"log"
	"net/http"
	"os"
	"simple-mux/database"
	"simple-mux/router"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	database.ConnectMySQL()
	router := router.InitRouter()
	log.Printf("Service is running on :%s", os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
