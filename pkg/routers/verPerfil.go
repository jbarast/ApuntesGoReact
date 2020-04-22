package routers

import (
	"PaginaWebGoReact/pkg/db"
	"encoding/json"
	"net/http"
)

func VerPerfil(w http.ResponseWriter, req *http.Request){
	ID := req.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := db.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro "+err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(perfil)
}
