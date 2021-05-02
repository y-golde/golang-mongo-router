package DB

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

func InitClient() (error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_URL")))
	
	fmt.Printf("API is " + os.Getenv("DB_URL") + "\n")
	if err != nil {return errors.New("cannot connect to mongo server" + err.Error())}
	err = Client.Ping(ctx, readpref.Primary())
	if err != nil {log.Fatal(err)}
	return nil
}