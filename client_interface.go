package kazooapi

import (
	"context"
	"net/http"
)

// APIClientInterface defines the methods that the APIClient should implement
type APIClientInterface interface {
	CallAPI(ctx context.Context, request *http.Request) (*http.Response, error)
	Authenticate(ctx context.Context) error
}
