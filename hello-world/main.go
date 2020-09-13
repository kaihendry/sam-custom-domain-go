package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/apex/gateway"
	"github.com/apex/log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gorilla/pat"
)

// pets database-ish
var pets = make(map[string]struct{})

func handler(lambdaRequest events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("got request", lambdaRequest)
	request, err := gateway.NewRequest(context.Background(), lambdaRequest)
	if err != nil {
		fmt.Println("newRequest error", err)
	}

	app := pat.New()
	app.Get("/", get)
	app.Post("/", post)

	res := httptest.NewRecorder()
	app.ServeHTTP(res, request)

	return events.APIGatewayProxyResponse{
		Body:       res.Body.String(),
		StatusCode: res.Code,
	}, nil
}

func main() {
	lambda.Start(handler)
}

// curl http://localhost:3000/
func get(w http.ResponseWriter, r *http.Request) {
	log.Info("list pets")

	if len(pets) == 0 {
		fmt.Fprintf(w, "no pets")
		return
	}

	for name := range pets {
		fmt.Fprintf(w, "- %s\n", name)
	}
}

// curl -d Tobi http://localhost:3000/
// curl -d Loki http://localhost:3000/
// curl -d Jane http://localhost:3000/
func post(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	name := string(b)
	pets[name] = struct{}{}
	log.WithField("name", name).Info("add pet")
	fmt.Fprintf(w, "welcome to the family %s!\n", name)
}
