package main

import (
	"PaginaWebGoReact/pkg/db"
	"PaginaWebGoReact/pkg/handlers"
	"log"
)

func main() {
	log.Println("Servicio inicializado")

	if db.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos.")
		return
	}

	handlers.Manejadores()

}
