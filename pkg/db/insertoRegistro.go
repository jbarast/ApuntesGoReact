package db

import (
	"PaginaWebGoReact/pkg/models"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)



func InsertoRegistro(usuario models.Usuario)(string, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), TiempoTimeoutSegundos*time.Second)
	defer cancel() // Cancelo el contexto WithTimeout.

	db := MongoConexion.Database(NombreBaseDatos)
	coleccion := db.Collection("usuarios")

	// Encripto la contraseña.
	usuario.Password, _ = EncriptarPassword(usuario.Password)

	result, err := coleccion.InsertOne(ctx, usuario)
	if err != nil{
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
