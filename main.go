package main

import (
	"context"
	"fmt"
	"log"

	"github.com/brockhenrie/api/controllers"
	"github.com/brockhenrie/api/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server *gin.Engine
	userservice services.UserService
	usercontroller	controllers.UserController
	ctx context.Context
	usercollection *mongo.Collection
	mongoclient *mongo.Client
	err error
)

func init(){
	ctx= context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://dev:Password123!!@localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil{
		log.Fatal(err)
	}
	err := mongoclient.Ping(ctx, readpref.Primary())
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("mongo connection has been established!")

	usercollection = mongoclient.Database("userdb").Collection("users")
	userservice = services.NewUserService(usercollection, ctx)
	usercontroller = controllers.New(userservice)

	server = gin.Default();
}

//v1/user/create
func main(){
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")

	usercontroller.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))
}