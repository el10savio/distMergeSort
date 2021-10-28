package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/el10savio/distMergeSort/sort"

	log "github.com/sirupsen/logrus"
)

// addBody is the format of the
// input JSON body to the
// Sort Handler
type addBody struct {
	Values []int `json:"values"`
}

// SortHandler is the HTTP handler used to sort in
// values to the sort node in the server
func SortHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody addBody

	// Obtain the values from POST Request Body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed parse request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Add the values into the Set
	list := sort.Sort(requestBody.Values)

	// DEBUG log in the case of success indicating
	// the new Set and the values added
	log.WithFields(log.Fields{
		"input_list":  requestBody.Values,
		"sorted_list": list,
	}).Debug("successful sort")

	// JSON encode response value
	json.NewEncoder(w).Encode(list)
}
