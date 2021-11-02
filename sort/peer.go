package sort

import (
	"encoding/json"
	"fmt"
	"math"
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
		return list, ErrNoPeers
	}

	sortedList := []int{}
	chunks := createChunks(list, int(math.Ceil(float64(len(list))/float64(len(peers)))))
	fmt.Println(chunks)

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
		return []int{}, ErrEmptyPeer
	}

	url := fmt.Sprintf("http://%s.%s/sort", peer, GetNetwork())
	JSONPayload, err := json.Marshal(Payload{Values: list})
	if err != nil {
		return []int{}, err
	}

	response, err := SendRequest(url, JSONPayload)
	if err != nil {
		return []int{}, err
	}

	values, err := processSortResponse(response)
	if err != nil {
		return []int{}, err
	}

	return values, nil
}

func processSortResponse(response *http.Response) ([]int, error) {
	var list []int
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []int{}, ErrInvalidHTTPStatusCode
	}

	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(&list)
	if err != nil {
		return []int{}, err
	}

	return list, nil
}
