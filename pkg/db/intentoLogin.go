package db

import (
	"PaginaWebGoReact/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func IntengoLogin(email string, password string) (models.Usuario, bool){
	usuario, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false{
		return usuario, false
	}

	passwordBytes := []byte(password) // No encriptada
	passwordDB := []byte(usuario.Password) // Encriptada

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return usuario, false
	}

	return usuario, true
}