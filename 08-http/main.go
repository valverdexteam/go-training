package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var (
	persons = []person{
		{
			Id:         1,
			FirstName:  "Pedro",
			LastName:   "Nascimento",
			InitialAge: 25,
		},
		{
			Id:         2,
			FirstName:  "Gerson",
			LastName:   "Souza",
			InitialAge: 35,
		},
		{
			Id:         4,
			FirstName:  "Lucas",
			LastName:   "Smith",
			InitialAge: 6,
		},
	}
)

type helloWordResponse struct {
	Message string
}

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/hello", helloWord)
	srv := setServer(handler)
	log.Printf("http server starting on port %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

func setServer(handler http.Handler) http.Server {
	return http.Server{
		Addr:         ":8000",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
}

func helloWord(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	response := &helloWordResponse{
		Message: "Hello Go Lang Word",
	}
	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(response)
}
