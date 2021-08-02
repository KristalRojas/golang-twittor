package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/KristalRojas/golang-twittor/middleware"
	"github.com/KristalRojas/golang-twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores Public
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleware.ChequeoBD(routers.Registro())).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
