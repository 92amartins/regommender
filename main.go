package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type recommendation struct {
	ID          string `json:"ID"`
	Source      string `json:"Source"`
	Target      string `json:"Target"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allRecommendations []recommendation

var recs = allRecommendations{
	{
		ID:          "1",
		Source:      "obama",
		Target:      "sapiens",
		Title:       "Sapiens: A brief history of humankind.",
		Description: "The book surveys the history of humankind from the evolution of archaic human species in the Stone Age up to the twenty-first century, focusing on Homo sapiens.",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createRecommendation(w http.ResponseWriter, r *http.Request) {
	var newRec recommendation
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Could not create recommendation")
	}

	json.Unmarshal(reqBody, &newRec)
	recs = append(recs, newRec)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newRec)
}

func getOneRec(w http.ResponseWriter, r *http.Request) {
	recID := mux.Vars(r)["id"]

	for _, singleRec := range recs {
		if singleRec.ID == recID {
			json.NewEncoder(w).Encode(singleRec)
		}
	}
}

func getAllRecs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(recs)
}

func deleteRec(w http.ResponseWriter, r *http.Request) {
	recID := mux.Vars(r)["id"]

	for i, singleRec := range recs {
		if singleRec.ID == recID {
			recs = append(recs[:i], recs[i+1:]...)
			fmt.Fprintf(w, "The recommendation %v has been deleted successfully", recID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)
	router.HandleFunc("/recommendation", createRecommendation).Methods("POST")
	router.HandleFunc("/recommendations/{id}", getOneRec).Methods("GET")
	router.HandleFunc("/recommendations/", getAllRecs).Methods("GET")
	router.HandleFunc("/recommendations/{id}", deleteRec).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
