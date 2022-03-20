package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"example.com/m/v2/controller"
	"example.com/m/v2/models"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ := mongo.Connect(ctx, clientOptions)
	models.Collection = client.Database("MongoDbGoland").Collection("Users")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controller.HomeLink)
	router.HandleFunc("/people", controller.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", controller.GetPersonByID).Methods("GET")
	router.HandleFunc("/people", controller.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", controller.UpdatePerson).Methods("PUT")
	router.HandleFunc("/people/{id}", controller.DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}