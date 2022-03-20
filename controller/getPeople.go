package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"example.com/m/v2/models"
	"go.mongodb.org/mongo-driver/bson"
)	

func GetPeople(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var people []models.Person
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	cursor, err := models.Collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person models.Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(people)
}
