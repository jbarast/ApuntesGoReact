package jwt

import (
	"PaginaWebGoReact/pkg/models"
	"log"
	"time"
	jwt "github.com/dgrijalva/jwt-go" // TODO: Mirar otra libreria.
)

func GeneroJWT(usuario models.Usuario) (string, error){

	miClave := []byte("Frase de seguridad") // Mirar como mejoarar esta parte

	payload := jwt.MapClaims{
		"email":     usuario.Email,
		"name":      usuario.Name,
		"fname":     usuario.FSubname,
		"lname":     usuario.LSubname,
		"birthday":  usuario.BirthDate,
		"biografia": usuario.Biografia,
		"ubicacion": usuario.Ubicacion,
		"sitioweb":  usuario.SitioWeb,
		"_id":       usuario.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil{
		log.Fatal("Error en la creacion del token jwt. Error: "+err.Error())
		return tokenStr, err
	}

	return tokenStr, nil
}