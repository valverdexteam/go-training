package main

import (
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

func main() {
	handler := http.NewServeMux()

	handler.HandleFunc("/hello", helloWord)

	srv := &http.Server{
		Addr:         ":8000",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	log.Printf("http server starting on port %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

func helloWord(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write([]byte(`{"message":"Hello World!"}`))
}
