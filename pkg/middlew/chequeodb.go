package middlew

import (
	"PaginaWebGoReact/pkg/db"
	"net/http"
)

// Funciones intermedias. "Un pasamanos"

func ChequeoDB(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){ //Funcion anonima.
		if db.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la base de datos",http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, req)
	}
}