package kazooapi

import (
	"net/http"
	"time"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")
)

// BasicAuth provides basic authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Realm    string `json:"realm,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
/*type APIKey struct {
	Key string `json:"api_key,omitempty"`
}

//SetAPIKey returns APIKey structure with the set key
func SetAPIKey(key string) (apiKey *APIKey) {
	return &APIKey{
		Key: key,
	}
}
*/
//Configuration is the structure which represents a configureation for APIClient
//you have to use either APIKey OR BasicAuth (meaning usernameme/password/realm)
//if you specify them both, then APIKey method will be chosen
type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	APIKey        string            `json:"api_key,omitempty"`
	BasicAuth     BasicAuth
	HTTPClient    *http.Client
	ClientTimeout time.Duration
}

//NewConfiguration prepares a default configuration structure for an APIClient
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      apiURL,
		DefaultHeader: make(map[string]string),
		UserAgent:     "kazoo-go/" + ver,
	}
	return cfg
}

//AddDefaultHeader is the method which easily adds necessary headers
//to the DefaultHeader map
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
