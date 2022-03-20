package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"example.com/m/v2/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPersonByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person models.Person
	id, _ := primitive.ObjectIDFromHex(mux.Vars(request)["id"])
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := models.Collection.FindOne(ctx, models.Person{ID: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(person)
}
