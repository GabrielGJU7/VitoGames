package main

import (
	"log"

	"github.com/GabrielGJU7/VitoGames/bd"
	"github.com/GabrielGJU7/VitoGames/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}
