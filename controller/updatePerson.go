package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"example.com/m/v2/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func UpdatePerson(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person models.Person
	id, _ := primitive.ObjectIDFromHex(mux.Vars(request)["id"])
	json.NewDecoder(request.Body).Decode(&person)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := models.Collection.UpdateOne(ctx, models.Person{ID: id}, bson.M{"$set": person})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}
