package routers

import (
	"PaginaWebGoReact/pkg/db"
	"PaginaWebGoReact/pkg/models"
	"encoding/json"
	"net/http"
)

func Registro(w http.ResponseWriter, req *http.Request){

	var usuario models.Usuario
	err := json.NewDecoder(req.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), http.StatusBadRequest) // 400
		return
	}

	if len(usuario.Email)==0 {
		http.Error(w, "El email de usuario es requerido", http.StatusBadRequest)
		return
	}

	if len(usuario.Password) <6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", http.StatusBadRequest)
		return
	}

	_, encontrado, _  := db.ChequeoYaExisteUsuario(usuario.Email)
	if encontrado {
		http.Error(w, "Ya existe un usuario ya registrado", http.StatusBadRequest)
		return
	}


	_, status, err := db.InsertoRegistro(usuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false{
		http.Error(w, "No se ha logrado insertar el registro del usuario", http.StatusBadRequest)
		return
	}


	w.WriteHeader(http.StatusCreated)
}