package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var users *mongo.Collection

func Connect(host, db string) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(host))
	if err != nil {
		return err
	}
	log.Printf("connected to mongodb host : %s", host)
	users = client.Database(db).Collection("users")
	return err
}
