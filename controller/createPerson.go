package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"example.com/m/v2/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePerson(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person models.Person
	json.NewDecoder(request.Body).Decode(&person)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := models.Collection.InsertOne(ctx, person)
	id := result.InsertedID
	person.ID = id.(primitive.ObjectID)
	json.NewEncoder(response).Encode(person)
}
