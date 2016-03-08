package endpoint

import (
	"errors"

	kitendpoint "github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

var (
	// ErrTokenVerificationFailure signifies that an auth token was missing or invalid
	ErrTokenVerificationFailure = errors.New("API token was missing or invalid")
)

// ProtectedOperationRequest is used to indicate a request type uses auth tokens
type ProtectedOperationRequest interface {
	GetToken() string
}

// VerifyAPIKey checks the API key in a request
func VerifyAPIKey(apiTokens []string) kitendpoint.Middleware {
	return func(next kitendpoint.Endpoint) kitendpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			message, ok := request.(ProtectedOperationRequest)
			if ok {
				suppliedAPIToken := message.GetToken()
				found := false
				for _, token := range apiTokens {
					if suppliedAPIToken == token {
						found = true
						break
					}
				}
				if !found {
					return nil, ErrTokenVerificationFailure
				}
			}
			return next(ctx, request)
		}
	}
}
