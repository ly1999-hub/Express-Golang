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

func DeletePerson(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	id, _ := primitive.ObjectIDFromHex(mux.Vars(request)["id"])
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := models.Collection.DeleteOne(ctx, &models.Person{ID: id})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}
