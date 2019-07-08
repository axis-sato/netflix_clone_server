package main

import (
	"encoding/json"
	"github.com/c8112002/netflix_clone_server/models"
	"log"
	"net/http"
)

func getItems(w http.ResponseWriter, r *http.Request)  {

	q := r.URL.Query().Get("query")

	res,err := json.Marshal(models.GetAllItems().Filter(q))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	http.HandleFunc("/items", getItems)
	http.Handle("/", http.FileServer(http.Dir("assets/images")))

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
