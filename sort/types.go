package sort

// Payload is the representation of the message
// passed between nodes to sort lists
type Payload struct {
	Values []int `json:"values"`
}
