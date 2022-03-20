package controller

import (
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection

func HomeLink(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Welcome home!")
}
