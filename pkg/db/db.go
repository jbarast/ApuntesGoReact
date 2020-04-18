package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConexion = ConectarDB()
const urlBaseDatos = "mongodb+srv://sa:sa@pruebas-wjylw.mongodb.net/test?retryWrites=true&w=majority"
var clienteOptions = options.Client().ApplyURI(urlBaseDatos)


/*ConectarDB es la funcion que me permite conectar la BBDD*/
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clienteOptions)
	if err != nil {
		log.Fatal("Error en creacicion de conexion:" + err.Error())
		return client // siempre tengo que devolver un client.
	}

	// Hacemos un ping, para ver si la base de datos esta levantada.
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa con base de datos de mongo.")
	return client
}


func ChequeoConnection() int {
	err := MongoConexion.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
