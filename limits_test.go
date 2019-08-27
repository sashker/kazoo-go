package kazooapi_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	kazooapi "gitlab.com/bmitelecom/kazoo-go"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestLimitsAPIService_GetLimits(t *testing.T) {
	ctx := context.Background()

	mux := http.NewServeMux()
	mux.HandleFunc("/v2/accounts/qe0ade400015367f0069d6dfbdca072a/limits", func(w http.ResponseWriter, r *http.Request){

		w.Header().Add("Server", "Cowboy")
		w.Header().Add("Content-Language", "en")
		w.Header().Add("Etag", "6a5ea7b838e7b91c5a38bc2b1050a4fa")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Vary", "accept-language, accept")
		w.Header().Add("Date", "Mon, 12 Aug 2019 05:18:39 GMT")

		body := `{
    "data": {
        "twoway_trunks": 1000,
        "inbound_trunks": 0,
        "allow_prepay": true,
        "outbound_trunks": 0,
        "id": "limits",
        "allow_postpay": false,
        "max_postpay_amount": 0
    },
    "revision": "2-01809fe4c9e8a85345215c0e29bdc2aa",
    "timestamp": "2019-08-20T11:34:28",
    "version": "4.2.33",
    "node": "o51LI7TmYyGMzigLbo-upw",
    "request_id": "8e97b5e8ee54c2e7057ed35e7480d08d",
    "status": "success",
    "auth_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjA4ZDQ0OTM1NDg0MTBhNDJkYjE1MWU1NTM0NjljODljIn0.eyJpc3MiOiJrYXpvbyIsImlkZW50aXR5X3NpZyI6Ikt3WFVMMHk0dWt6cWhiVWU0Y1Z2ZHNUb1Y4SVQ2NGpEMFdRc3Q4aVVucjgiLCJhY2NvdW50X2lkIjoiZmUwYWRlNDAwMDE1MzY3ZjAwNjlkNmRmYmRjYTA3MmEiLCJtZXRob2QiOiJjYl9hcGlfYXV0aCIsImV4cCI6MTU2NjMwNDM4M30.qUbrkyzzDO0wVyGpdsAE6OwI0BSA2tCMFB8xJ7JF3dpqW_j6PCMdiiKc_cnxQGCTblabnEeeOTrKbOB9KOQ3wULVto106Po6JJcXjz9VsPi_P77jdAefiTldYbBmXfwfT9wSA-XTMSHYrCpfeizvBnGUzF1XA7KnfhF1BkCMRqBeO-8kdZG7fuMyhX_XPr-VbP9sTFTbNdYxP9ZgqhkrLodODyH2gn1JFSvPY9ob_ixGGfeF3WSKZMMgv1onc9bztjOwcG31luNGYhXVGSEd4HScyidQ7bxlJnDuqyksKFdG2WTeBSvQ31FRq3fbiH3-MYaiNSmKN9WRt0b-pdgH7g"
}`
		io.WriteString(w, body)

	})
	mux.HandleFunc("/v2/api_auth", func(w http.ResponseWriter, r *http.Request){

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

	srv := httptest.NewServer(mux)
	defer srv.Close()


	cfg := kazooapi.NewConfiguration()
	cfg.APIKey = "e0a582bad3fb7fe3897ebf70cc0f542bbdc9a17895764266f094b953254d3d84"
	cfg.BasePath = srv.URL + "/v2"

	clt, err := kazooapi.NewAPIClient(cfg)
	if err != nil {
		t.Error("Can't create the API client")
	}

	//input := &kazooapi.Account{ID: "qe0ade400015367f0069d6dfbdca072a"}

	resp, err := clt.LimitsAPI.GetLimits(ctx, "qe0ade400015367f0069d6dfbdca072a")
	if err != nil {
		t.Error(err)
	}

	t.Logf("Response data %#v", resp)

	assert.Equal(t, int64(1000), resp.TwowayTrunks, "ID's should be equal" )
}