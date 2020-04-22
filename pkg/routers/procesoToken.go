package routers

import (
	"PaginaWebGoReact/pkg/db"
	"PaginaWebGoReact/pkg/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string)(*models.Claim, bool, string, error){
	miClave := []byte("Frase de seguridad")
	claims := &models.Claim{}

	// El token siempre empieza con la palabra Bearear que no es parte del token.
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token)(interface{}, error){
		return miClave, nil
	})
	if err != nil{
		return claims, false, string(""), err
	}

	if !tkn.Valid{
		return claims, false, string(""), errors.New("token invalido")
	}

	_, encontrado, _ := db.ChequeoYaExisteUsuario(claims.Email)
	if encontrado == true {
		Email = claims.Email
		IDUsuario = claims.ID.Hex()
	}
	return claims, encontrado, IDUsuario, nil
}
