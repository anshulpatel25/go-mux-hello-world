package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", IndexHandler).Methods("GET")
	v1Router := router.PathPrefix("/v1").Subrouter()

	v1Router.HandleFunc("/greet", GreetHandler).Methods("GET")
	return router
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, req)
	return response
}

func TestIndexEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(request)
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	actual := string(bodyBytes)
	expected := "Simple REST API which greets users"
	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestGreetEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/greet", nil)
	response := executeRequest(request)
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	actual := string(bodyBytes)
	expected := "Greetings user of"
	if !strings.Contains(actual, expected) {
		t.Errorf("Actual response %s doesn't contain %s", actual, expected)
	}
}
