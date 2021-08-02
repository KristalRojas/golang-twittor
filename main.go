package main

import (
	"log"

	"github.com/KristalRojas/golang-twittor/database"
	"github.com/KristalRojas/golang-twittor/handlers"
)

func main() {
	if database.ChequeoConexion() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}
