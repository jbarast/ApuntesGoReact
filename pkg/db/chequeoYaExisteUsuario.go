package db

import (
	"PaginaWebGoReact/pkg/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string){
	ctx, cancel := context.WithTimeout(context.Background(), TiempoTimeoutSegundos*time.Second)
	defer cancel()

	db := MongoConexion.Database(NombreBaseDatos)
	colecion := db.Collection("usuarios")

	// Hacemos una busqueda
	condicion := bson.M{"email":email}

	var resultado models.Usuario

	err := colecion.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex() // Convertito a string un objeto ID
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}