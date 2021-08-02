package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://AdminKristal:AkQ0eH6AnvLBCitu@proyecto-golang.gnm1m.mongodb.net/test")

//ConectarBD Public
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion existosa con la BD")
	return client
}

// ChequeoConexion Public
func ChequeoConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
