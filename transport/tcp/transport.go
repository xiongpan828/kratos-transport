package tcp

import (
	"context"
	"net/http"

	"github.com/go-kratos/kratos/v2/transport"
)

const (
	KindTcp transport.Kind = "tcp"
)

var _ Transporter = &Transport{}

type Transporter interface {
	transport.Transporter
	Request() *http.Request
	PathTemplate() string
}

// Transport is a tcp transport.
type Transport struct {
	endpoint  string
	operation string
}

// Kind returns the transport kind.
func (tr *Transport) Kind() transport.Kind {
	return KindTcp
}

// Endpoint returns the transport endpoint.
func (tr *Transport) Endpoint() string {
	return tr.endpoint
}

// Operation returns the transport operation.
func (tr *Transport) Operation() string {
	return tr.operation
}

// Request returns the HTTP request.
func (tr *Transport) Request() *http.Request {
	return nil
}

// RequestHeader returns the request header.
func (tr *Transport) RequestHeader() transport.Header {
	return nil
}

// ReplyHeader returns the reply header.
func (tr *Transport) ReplyHeader() transport.Header {
	return nil
}

// PathTemplate returns the http path template.
func (tr *Transport) PathTemplate() string {
	return ""
}

// SetOperation sets the transport operation.
func SetOperation(ctx context.Context, op string) {
	if tr, ok := transport.FromServerContext(ctx); ok {
		if tr, ok := tr.(*Transport); ok {
			tr.operation = op
		}
	}
}

type headerCarrier struct{}

// Get returns the value associated with the passed key.
func (hc headerCarrier) Get(key string) string {
	return ""
}

// Set stores the key-value pair.
func (hc headerCarrier) Set(key, value string) {
}

// Keys lists the keys stored in this carrier.
func (hc headerCarrier) Keys() []string {
	return []string{}
}
