// grpc-like service

package protocol

type RequestType int

const (
	Subscribe RequestType = iota
	Unsubscribe
)

// Request is a wrapper around all requests
type Request struct {
	RequestType RequestType
	data        any
}

// SubscribeRequest is used to tell which topic a publisher is willing to subscribe to
type SubscribeRequest struct {
	PublisherID int
	Topic       string
}

// SubscribeResponse is a response to SubscribeRequest
type SubscribeResponse struct {
	TopicMetadata any
}

// UnsubscribeRequest is used to tell which topic a publisher is willing to unsubscribe from
type UnsubscribeRequest struct{}

// UnsubscribeResponse is a response to UnsubscribeRequest
type UnsubscribeResponse struct{}

type Protocol interface {
	Subscribe(SubscribeRequest) (SubscribeResponse, error)
	Unsubscribe(UnsubscribeRequest) (UnsubscribeResponse, error)
}
