package workerpool

import "time"

// Request represent a request to be processed by a worker
type Request struct {
	Handler    RequestHandler
	Type       int
	Data       interface{}
	Timeout    time.Duration
	Retries    int
	MaxRetries int
}

// RequestHandler defines a function type for handling requests
type RequestHandler func(interface{}) error
