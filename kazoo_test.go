package kazooapi_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	kazooapi "gitlab.com/bmitelecom/kazoo-go"
)

func MockKazooServer(t *testing.T, body string) *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/v2/api_auth", func(w http.ResponseWriter, r *http.Request) {
		t.Log("Hit api_auth")

		w.Header().Add("Server", "Cowboy")
		w.Header().Add("Content-Language", "en")
		w.Header().Add("Etag", "6a5ea7b838e7b91c5a38bc2b1050a4fa")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Vary", "accept-language, accept")
		w.Header().Add("Date", "Mon, 12 Aug 2019 05:18:39 GMT")

		w.WriteHeader(201)

		body := `{
    "page_size": 1,
    "data": {
        "account_id": "fe0ade400015367f0069d6dfbdca072a",
        "reseller_id": "fe0ade400015367f0069d6dfbdca072a",
        "is_reseller": true,
        "account_name": "root",
        "language": "en-US",
        "apps": [
            {
                "id": "1d3b135195e3c86abe8ed85e9d2a5210",
                "name": "accounts",
                "api_url": "https://api.pbx.example.com/v2",
                "label": "Accounts Manager"
            },
            {
                "id": "a2eeff7655a9e3b86c2ed9916a3f2a34",
                "name": "admin",
                "label": "Admin Manager"
            },
            {
                "id": "12556b35127b8f2819d4b1f67e03baef",
                "name": "callflows",
                "api_url": "https://api.pbx.example.com/v2",
                "label": "Callflows"
            },
            {
                "id": "c10c209fd993e33a7d77f21716787ca2",
                "name": "fax",
                "api_url": "https://api.pbx.example.com/v2",
                "label": "Fax Manager"
            },
            {
                "id": "f6e0ecda13b18411dcdb31142cee94bf",
                "name": "numbers",
                "api_url": "https://api.pbx.example.com/v2",
                "label": "Number Manager"
            },
            {
                "id": "559d91f140f192bb15982952ff59f7ff",
                "name": "pbxs",
                "api_url": "https://api.pbx.example.com/v2",
                "label": "PBX Connector"
            },
            {
                "id": "f33ab81f003fcf57c0e729570e6f8f89",
                "name": "voicemails",
                "api_url": "https://api.pbx.example.com/v2",
                "label": "Voicemail Manager"
            },
            {
                "id": "ac21058f2f5ac5bc7ea71e5c0fb7cf1a",
                "name": "voip",
                "api_url": "https://api.pbx.example.com/v2",
                "label": "Smart PBX"
            },
            {
                "id": "948102dbcb8d311730a8f64c32dce944",
                "name": "webhooks",
                "api_url": "https://api.pbx.example.com/v2",
                "label": "Webhook Manager"
            }
        ]
    },
    "revision": "automatic",
    "timestamp": "2019-08-17T05:41:45",
    "version": "4.2.33",
    "node": "o51LI7TmYyGMzigLbo-upw",
    "request_id": "84df73edb39079dda1c1149c642edfab",
    "status": "success",
    "auth_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjA4ZDQ0OTM1NDg0MTBhNDJkYjE1MWU1NTM0NjljODljIn0.eyJpc3MiOiJrYXpvbyIsImlkZW50aXR5X3NpZyI6Ikt3WFVMMHk0dWt6cWhiVWU0Y1Z2ZHNUb1Y4SVQ2NGpEMFdRc3Q4aVVucjgiLCJhY2NvdW50X2lkIjoiZmUwYWRlNDAwMDE1MzY3ZjAwNjlkNmRmYmRjYTA3MmEiLCJtZXRob2QiOiJjYl9hcGlfYXV0aCIsImV4cCI6MTU2NjAyNDEwNX0.d2CP5F7XP4QbblrCcRlfDBXxjreiOX_K6X34GPo_tzG-ct0S-GhHjxpaBAPm7YY1HmHAODaso94wEm1zzJK-tzZVrAcvJCKXD2pxR7FCUD76XL_ZJh2zXIxXbf6lxBMejKILBQvw1Oq3yPihyN1WPa61RKuPkXqbbGL1cnUesx6RqFlD6tvRxMIevV27U1y6UaYqM5JyT1u37HseCyM5XAvssMwlTFkbrdh0Oh4gs1w0SeYS63_H5IeeqLmx-QWmjmemCbh5jNAWECu1ZWxyisqhDx2GhN8SmXh1qZ2qBABhSV4O1_f-hpYorNZcr5tLbMjxBG9J9Br_tbsqvQmAxQ"
}`
		io.WriteString(w, body)
	})

	mux.HandleFunc("/v2/accounts", func(w http.ResponseWriter, r *http.Request) {
		t.Log("Hit accounts")

		w.Header().Add("Server", "Cowboy")
		w.Header().Add("Content-Language", "en")
		w.Header().Add("Etag", "6a5ea7b838e7b91c5a38bc2b1050a4fa")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Vary", "accept-language, accept")
		w.Header().Add("Date", "Mon, 12 Aug 2019 05:18:39 GMT")

		io.WriteString(w, body)
	})

	mux.HandleFunc("/v2/accounts/qe0ade400015367f0069d6dfbdca072a/phone_numbers", func(w http.ResponseWriter, r *http.Request) {
		t.Log("Hit phone_numbers")

		w.Header().Add("Server", "Cowboy")
		w.Header().Add("Content-Language", "en")
		w.Header().Add("Etag", "6a5ea7b838e7b91c5a38bc2b1050a4fa")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Vary", "accept-language, accept")
		w.Header().Add("Date", "Mon, 12 Aug 2019 05:18:39 GMT")

		if r.Method == "DELETE" {
			body = `{
				"data": {
					"message": "bad identifier",
					"not_found": "The number could not be found"
				},
				"error": "404",
				"message": "bad_identifier",
				"status": "error",
				"timestamp": "2020-03-03T07:10:35",
				"version": "4.2.33",
				"node": "o51LI7TmYyGMzigLbo-upw",
				"request_id": "be0631b976148f384f2d1310b21df3bd",
				"auth_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjA4ZDQ0OTM1NDg0MTBhNDJkYjE1MWU1NTM0NjljODljIn0.eyJpc3MiOiJrYXpvbyIsImlkZW50aXR5X3NpZyI6Ikt3WFVMMHk0dWt6cWhiVWU0Y1Z2ZHNUb1Y4SVQ2NGpEMFdRc3Q4aVVucjgiLCJhY2NvdW50X2lkIjoiZmUwYWRlNDAwMDE1MzY3ZjAwNjlkNmRmYmRjYTA3MmEiLCJtZXRob2QiOiJjYl9hcGlfYXV0aCIsImV4cCI6MTU4MzIyMzAwNn0.oObFGrAhMNp9YhyK-0mk9nXpU4Stl-QJAFOOreMQVUT5flZU1qDRBU_l6RK2tYWYU8gSqOTsZx3FQaAbMLbbhMsaSUNNIJdXpkv-vz9FpabzGKIvMILkTAJePZF-JlePsBguDvt9bkODnLi_tXduygb5-cHRAF-HA757ovnD_8mS1aws7dxofJlfTeyFDjq8JVH60oZ1fBqhtG_smEPpWbyzbxspNp56RRqxWY5ts568qOPoHdIYB-SbpwCZgGcnKzAjRnYeJ2Ae2yACRcITmL6TeT2dSfWLd8asVm-szo1RBQaxNIYM7t_TFXm_qL2TRWH1CfB1zfeVzvH9eQlYxA"
			}`

			w.WriteHeader(404)
		}

		io.WriteString(w, body)
	})

	srv := httptest.NewServer(mux)

	return srv
}

func TestMakingNewAPIClient(t *testing.T) {
	cfg := kazooapi.NewConfiguration()

	t.Logf("Configuration looks like %#v", cfg)

	_, err := kazooapi.NewAPIClient(cfg)
	if err != nil {
		assert.Equal(t, "you have to specify APIKey or username/password/realm", err.Error())
	}

	cfg.APIKey = "abcdef123456"
	cfg.BasicAuth = kazooapi.BasicAuth{Username: "user", Password: "pass", Realm: "test"}

	_, err2 := kazooapi.NewAPIClient(cfg)
	if err2 != nil {
		t.Error(err)
	}

}
