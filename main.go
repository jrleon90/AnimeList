package main

import (
	"encoding/json"
	api "example/animeFactAPI/repository"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/anime/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint Hit: All Articles Endpoint")
		vars := mux.Vars(r)
		key := vars["id"]
		anime, err := api.GetAnimeFact(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(anime.Data)
	})
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()

}
