package main

import (
	"context"
	"fmt"
	"log"

	"github.com/JoseVilledaa/superheroes-api/controllers"
	"github.com/JoseVilledaa/superheroes-api/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	ss          services.SuperheroeService
	sc          controllers.SuperheroeController
	ctx         context.Context
	superheroec *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connected to MongoDB!")

	superheroec = mongoclient.Database("superheroes").Collection("superheroes")
	ss = services.NewSuperheroeService(superheroec, ctx)
	sc = controllers.New(ss)
	server = gin.Default()
}

func main() {
	basepath := server.Group("/v1")
	sc.RegisterRoutes(basepath)
	log.Fatal(server.Run(":8080"))
}
