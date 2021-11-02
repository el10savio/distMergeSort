package sort

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// GetPeerList Obtains Peer List
// From Environment Variable
func GetPeerList() []string {
	return strings.Split(os.Getenv("PEERS"), ",")
}

// GetHost Obtains Host
// From Environment Variable
func GetHost() string {
	return os.Getenv("HOST")
}

// GetPeerListWithoutHost Obtains Peer List
// From Environment Variable
// Excluding The Host
func GetPeerListWithoutHost() []string {
	peers, host := GetPeerList(), GetHost()
	peerList := make([]string, 0)

	for _, peer := range peers {
		if peer != host {
			peerList = append(peerList, peer)
		}
	}

	return peerList
}

// GetNetwork Obtains Network
// From Environment Variable
func GetNetwork() string {
	return os.Getenv("NETWORK") + ":8080"
}

// SendRequest handles sending of an HTTP POST Request
func SendRequest(url string, payload []byte) (*http.Response, error) {
	if url == "" {
		return nil, ErrEmptyURL
	}

	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		err = fmt.Errorf("failed to create POST request: %v", err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		err = fmt.Errorf("failed to send POST request: %v", err)
		return nil, err
	}

	return response, nil
}
