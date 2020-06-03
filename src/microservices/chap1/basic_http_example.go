package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type validationContextKey string

type helloRequest struct {
	Name string `json:"name"`
}

type helloResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080
	handler := newValidRequestHandler(newHelloRequestHandler())
	http.Handle("/hello", handler)
	log.Printf("Server running on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

type validRequestHandler struct {
	next http.Handler
}

func newValidRequestHandler(next http.Handler) http.Handler {
	return validRequestHandler{next: next}
}

func (h validRequestHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r = r.WithContext(c)

	h.next.ServeHTTP(rw, r)
}

type helloRequestHandler struct{}

func newHelloRequestHandler() http.Handler {
	return helloRequestHandler{}
}

func (h helloRequestHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(validationContextKey("name")).(string)
	response := helloResponse{Message: "hello " + name}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)

}
