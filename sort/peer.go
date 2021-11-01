package sort

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var peers []string

func init() {
	peers = GetPeerListWithoutHost()
}

// peerSort splits the list into chunks
// and send's a chunk to each peer
func peerSort(list []int) ([]int, error) {
	if len(peers) == 0 {
		return list, errors.New("no peers available")
	}

	sortedList := []int{}
	chunks := createChunks(list, len(peers))

	// For Each Chunk, Send Peer Sort Request
	for index, chunk := range chunks {
		chunkSorted, err := sendSortRequest(chunk, peers[index])
		if err != nil {
			return list, err
		}
		sortedList = merge(sortedList, chunkSorted)
	}

	return sortedList, nil
}

func createChunks(slice []int, chunkSize int) [][]int {
	if len(slice) == 0 || chunkSize == 0 {
		return [][]int{}
	}

	if len(slice) < chunkSize {
		chunkSize = len(slice)
	}

	return append(
		createChunks(slice[chunkSize:], chunkSize),
		slice[0:chunkSize],
	)
}

// sendSortRequest sends the HTTP Sort
// POST request to the given peer
func sendSortRequest(list []int, peer string) ([]int, error) {
	if peer == "" {
		return []int{}, errors.New("empty peer provided")
	}

	url := fmt.Sprintf("http://%s.%s/sort", peer, GetNetwork())
	JSONPayload, err := json.Marshal(Payload{list})
	if err != nil {
		return []int{}, err
	}

	response, err := SendRequest(url, JSONPayload)
	if err != nil {
		return []int{}, err
	}

	if response.StatusCode != http.StatusOK {
		return []int{}, errors.New("received invalid response code")
	}

	var responseBody Payload

	// Obtain the values from POST Request Body
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&responseBody)
	if err != nil {
		return []int{}, err
	}

	return responseBody.Values, nil
}
