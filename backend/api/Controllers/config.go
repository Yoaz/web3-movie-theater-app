package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Server type to pass around contains mongoDB & mux Router instance
type Server struct {
	DB *mongo.Database
	Router *mux.Router
}

//Initialize Server struct's MongoDB database & mux Router instances
func (server *Server) InitDBRouter(connectionString, dbName string){
	// Context with timeout declaration
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	// Client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect
	client, err := mongo.Connect(ctx, clientOption)

	// Connection fails
	if err != nil {
		log.Fatalf("There was an error with the MongoDB connection %s", err)
	}

	// Based on Mongo documentation recommendation
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()
	
	// Call Ping to verify that the deployment is up and the Client was
	// configured successfully. As mentioned in the Ping documentation, this
	// reduces application resiliency as the server may be temporarily
	// unavailable when Ping is called.
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	//Assign values to type fields
	server.Router = mux.NewRouter()
	server.DB = client.Database(dbName)
	server.InitRoutes()

	// Connection success
	fmt.Println(("MongoDB connection success!"))
}

// func (server *Server) establishCollections() error {
// 	server.DB.CreateCollection()
// }

// Run server
func (server *Server) RunServer(addr, port string){
	fmt.Printf("Listening to port %s\n", port)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}