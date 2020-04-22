package db

import (
	"PaginaWebGoReact/pkg/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func BuscoPerfil(ID string)(models.Usuario, error){
	ctx, cancel := context.WithTimeout(context.Background(),TiempoTimeoutSegundos*time.Second)
	defer cancel()


	db := MongoConexion.Database(NombreBaseDatos)
	coleccion := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := coleccion.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil{
		fmt.Println("Registro no encontrado "+err.Error())
		return perfil, err
	}


	return perfil, nil

}
