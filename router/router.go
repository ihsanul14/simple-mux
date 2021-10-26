package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	mod1_router "simple-mux/router/modul1"
)

//InitRouter is used for initiate router
func InitRouter() http.Handler {
	router := mux.NewRouter()

	mod1_router.Router(router)
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
	}).Handler(router)

	return handler
}
