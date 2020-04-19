package routers

import (
	"PaginaWebGoReact/pkg/db"
	"PaginaWebGoReact/pkg/jwt"
	"PaginaWebGoReact/pkg/models"
	"encoding/json"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, req *http.Request){
	w.Header().Add("content-type","application/json") // TODO: Mirar si se recomendia poner algo mas.

	var usuario models.Usuario

	err := json.NewDecoder(req.Body).Decode(&usuario)
	if err != nil{
		http.Error(w, "Usuario y/o Contraseña invalidos" + err.Error(), http.StatusBadRequest)
		return
	}

	if len(usuario.Email) == 0 {
		http.Error(w, "El email del usuario es requerido "+err.Error(), http.StatusBadRequest)
		return
	}

	documento, existe := db.IntengoLogin(usuario.Email, usuario.Password)
	if existe == false{
		http.Error(w, "Usuario y/o contraseña incorrectos", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil{
		http.Error(w, "Ocurrio un error al intentar generar el Toker correspondiente " + err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.RespuestaLogin {
		Token : jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)


	// Si queremos grabarla en una coocky.
	expirationTime := time.Now().Add(24 * time.Hour) // Fecha limite que durara el token de key
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})
}
