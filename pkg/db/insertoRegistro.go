package db

import (
	"PaginaWebGoReact/pkg/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)


func InsertoRegistro(usuario models.Usuario)(string, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), TiempoTimeoutSegundos*time.Second)
	defer cancel() // Cancelo el contexto WithTimeout.

	db := MongoConexion.Database(NombreBaseDatos)
	coleccion := db.Collection("usuarios")

	// Encripto la contraseña.
	usuario.Password, _ = EncriptarPassword(usuario.Password)

	//Todo: Mejorar codigo para que no necesite crear un bson sin el objeto id.
	usuarioDoc := bson.D{
		{"name",usuario.Name},
		{"fsubname",usuario.FSubname},
		{"lsubname", usuario.LSubname },
		{"birthdate", usuario.BirthDate},
		{"email", usuario.Email},
		{"password", usuario.Password},
		{"avatar", usuario.Avatar},
		{"banner", usuario.Banner},
		{"biografia" , usuario.Biografia},
		{"ubicacion" , usuario.Ubicacion},
		{"sitioweb", usuario.SitioWeb},
	}

	result, err := coleccion.InsertOne(ctx, usuarioDoc)
	if err != nil{
		log.Println("Error al insertar en la base de datos "+err.Error())
		return "", false, err
	}


	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
