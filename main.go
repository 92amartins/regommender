package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

type recommendation struct {
	Source string  `json:"Source"`
	Target string  `json:"Target"`
	Score  float64 `json:"Score"`
}

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is reGommender. I'm up and running.")
}

func getRecommendation(w http.ResponseWriter, r *http.Request) {
	source := mux.Vars(r)["source_id"]
	response, _ := client.ZRevRange(source, 0, 9).Result()
	json.NewEncoder(w).Encode(response)
}

func setRecommendation(w http.ResponseWriter, r *http.Request) {
	var rec recommendation
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &rec)
	client.ZAdd(rec.Source, redis.Z{
		Score:  rec.Score,
		Member: rec.Target,
	})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", healthcheck).Methods("GET")
	router.HandleFunc("/recommendations/{source_id}", getRecommendation).Methods("GET")
	router.HandleFunc("/recommendation/", setRecommendation).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
