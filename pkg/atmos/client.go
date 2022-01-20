package atmos

import (
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/umatare5/atmos-go"
)

// NewClient ...
func NewClient(server string, opts ...atmos.ClientOption) (*atmos.ClientWithResponses, error) {
	return atmos.NewClientWithResponses(server)
}

// NewClientWithToken ...
func NewClientWithToken(server string, token string) (*atmos.ClientWithResponses, error) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken(token)
	if err != nil {
		panic(err)
	}

	return atmos.NewClientWithResponses(
		server, atmos.WithRequestEditorFn(bearerTokenProvider.Intercept),
	)
}
