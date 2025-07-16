package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Println(err)
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// validation

	if reqBody.UserID == "" {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}

	// TODO: Call trip service
	writeJSON(w, http.StatusCreated, "ok")

}
