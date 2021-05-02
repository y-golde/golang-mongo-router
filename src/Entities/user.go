package Entities

type User struct {
	Id         string `json:"id" bson:"id"`
	First_name string `json:"firstName" bson:"first_name"`
	Last_name  string `json:"lastName" bson:"last_name"`
}