package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Basic Go api application.",
		Description: "Test application.",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home buddy!")
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "You haven't made a correct request.")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)

}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent event
	var singleEventID string

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Something went wrong while attemtping to read the update event request body")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	// now go through the even table and find relevant event
	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			events = append(events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
			singleEventID = singleEvent.ID
		}
	}
	if len(singleEventID) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No record found with id => "+eventID)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent)
	router.HandleFunc("/event/{id}", getOneEvent)
	router.HandleFunc("/events", getAllEvents)
	router.HandleFunc("/update/{id}", updateEvent)
	log.Fatal(http.ListenAndServe(":8080", router))
}
