package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/el10savio/distMergeSort/sort"
	log "github.com/sirupsen/logrus"
)

// SortHandler is the HTTP handler used to sort in
// values to the sort node in the server
func SortHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody sort.Payload

	// Obtain the values from POST Request Body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed parse request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DEBUG log to print input list
	log.WithFields(log.Fields{
		"input_list": requestBody.Values,
	}).Debug("received list")

	// Sort the values
	list, err := sort.Sort(requestBody.Values)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to sort values")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DEBUG log in the case of success
	// indicating the new sorted list
	log.WithFields(log.Fields{
		"sorted_list": list,
	}).Debug("successful sort")

	// JSON encode response value
	json.NewEncoder(w).Encode(list)
}
