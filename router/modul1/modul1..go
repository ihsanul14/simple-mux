package modul1

import (
	handler "simple-mux/handler/modul1"

	"github.com/gorilla/mux"
)

func Router(router *mux.Router) {
	router.HandleFunc("/api/data", handler.DataHandler).Methods("POST")
	// router.HandleFunc("api/data/add", handler.CreateHandler)
	// router.HandleFunc("api/data/update", handler.UpdateHandler)
	// router.HandleFunc("api/data/delete", handler.DeleteHandler)
}
