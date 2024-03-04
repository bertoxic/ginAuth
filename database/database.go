package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)




func DBinstance() *mongo.Client {
   err := godotenv.Load(".env")
   if err != nil {
    log.Fatal("error loading .env file")
   }
    MongoDb := os.Getenv("MONGODB_URL")

    client , err :=mongo.NewClient(options.Client().ApplyURI(MongoDb))
    if err != nil {
        log.Fatal(err)
    }
    ctx,  cancel:= context.WithTimeout(context.Background(), 14* time.Second)
    defer cancel()
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("connected to mongodb!")

    return client
   }
   
   var Client *mongo.Client = DBinstance()

   func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{
    var Collection *mongo.Collection = client.Database("cluster").Collection(collectionName)

    return Collection
   }