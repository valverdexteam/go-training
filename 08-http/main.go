package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ErrorResponse struct {
	Message string
}

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/persons", handlePersons)
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

// I know that this is not the ideal solution but I'm just trying to play around with the API
// For the final project I'll pick up one http framework to build a real rest API
func handlePersons(responseWriter http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	if request.Method == http.MethodGet {
		id, present := query["id"]
		if !present || len(id) == 0 {
			fmt.Println("Person id is not present")
			getAllPersons(responseWriter, request)
			return
		}
		getPerson(responseWriter, request, id[0])
		return
	}

	if request.Method == http.MethodPost {
		updatePerson(responseWriter, request)
	}

	if request.Method == http.MethodPut {
		addPerson(responseWriter, request)
	}

	if request.Method == http.MethodDelete {
		id, present := query["id"]
		if !present || len(id) == 0 {
			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(ErrorResponse{"Id is required for deletion"})
			responseWriter.WriteHeader(http.StatusBadRequest)
			return
		}
		deletePerson(responseWriter, request, id[0])
	}
}

func getAllPersons(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	persons := listAllPersons()
	if len(persons) == 0 {
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(persons)
}

func getPerson(responseWriter http.ResponseWriter, request *http.Request, id string) {
	intId, conversionError := strconv.Atoi(id)
	if conversionError != nil {
		responseWriter.Header().Set("Content-Type", "application/json")
		json.NewEncoder(responseWriter).Encode(ErrorResponse{conversionError.Error()})
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	person, error := findById(intId)
	if error != nil {
		responseWriter.WriteHeader(http.StatusNotFound)
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(person)
}

func updatePerson(responseWriter http.ResponseWriter, request *http.Request) {
	var payload person
	json.NewDecoder(request.Body).Decode(&payload)

	updateError := update(payload)
	if updateError != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Header().Set("Content-Type", "application/json")
		json.NewEncoder(responseWriter).Encode(ErrorResponse{updateError.Error()})
	}
	responseWriter.WriteHeader(http.StatusOK)
}

func addPerson(responseWriter http.ResponseWriter, request *http.Request) {
	var payload person
	json.NewDecoder(request.Body).Decode(&payload)

	insertError := update(payload)
	if insertError != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Header().Set("Content-Type", "application/json")
		json.NewEncoder(responseWriter).Encode(ErrorResponse{insertError.Error()})
	}
	responseWriter.WriteHeader(http.StatusOK)
}

func deletePerson(responseWriter http.ResponseWriter, request *http.Request, id string) {
	intId, conversionError := strconv.Atoi(id)
	if conversionError != nil {
		responseWriter.Header().Set("Content-Type", "application/json")
		json.NewEncoder(responseWriter).Encode(ErrorResponse{conversionError.Error()})
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	error := delete(intId)
	if error != nil {
		responseWriter.Header().Set("Content-Type", "application/json")
		json.NewEncoder(responseWriter).Encode(ErrorResponse{error.Error()})
		responseWriter.WriteHeader(http.StatusNotFound)
		return
	}
	responseWriter.WriteHeader(http.StatusOK)
}
