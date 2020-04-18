package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)


type Usuario struct {
	ID primitive.ObjectID `bson:"_id, omitempty" json:"id"`  // Debemos indicar el tipo de dato.
	Name string `bson:"name" json:"name"`
	FSubname string `bson:"fsubname" json:"fSubname"`
	LSubname string `bson:"lsubname" json:"lSubname"`
	BirthDate time.Time `bson:"birthdate" json:"birthDate"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password,omitempty"` // Nunca devolver una contraseña.
	Avatar string `bson:"avatar" json:"avatar,omitempty"`
	Banner string `bson:"banner" json:"banner,omitempty"`
	Biografia string `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion string `bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb string `bson:"sitioweb" json:"sitioWeb,omitempty"`
}
