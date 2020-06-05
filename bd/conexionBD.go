package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto a la coonexion a la BD*/
var MongoCN = ConectarBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://vito:admin@vitogames-35a04.mongodb.net/test?retryWrites=true&w=majority")

/*ConectarBD es la funcion que conecta la bd*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa")
	return client
}

/*ChequeoConnection es el Ping a la BD */
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
