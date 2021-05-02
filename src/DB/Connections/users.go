package Connections

import (
	"helloworld/src/DB"

	"go.mongodb.org/mongo-driver/mongo"
)


func GetUsersCollection() (*mongo.Collection) {
	return DB.Client.Database("router-test").Collection("users")
}