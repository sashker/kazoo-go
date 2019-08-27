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


func TestCallflowsAPIService_ListCallflowsAPIService(t *testing.T) {
	ctx := context.Background()

	mux := http.NewServeMux()
	mux.HandleFunc("/v2/accounts/qe0ade400015367f0069d6dfbdca072a/callflows", func(w http.ResponseWriter, r *http.Request){

		w.Header().Add("Server", "Cowboy")
		w.Header().Add("Content-Language", "en")
		w.Header().Add("Etag", "6a5ea7b838e7b91c5a38bc2b1050a4fa")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Vary", "accept-language, accept")
		w.Header().Add("Date", "Mon, 12 Aug 2019 05:18:39 GMT")

		body := `{
    "next_start_key": "g20AAAAgOWVlZjA4NzhlN2JhYjI4OTJjYTIxNTYxZjA1OTFjNjY",
    "page_size": 50,
    "data": [
        {
            "id": "9e1e5f9031e9e8446f54da9df47680a0",
            "numbers": [],
            "patterns": [
                "\\*72([0-9]*)$"
            ],
            "featurecode": {
                "name": "call_forward[action=activate]",
                "number": "72"
            }
        },
        {
            "id": "970af9b5ad818ffc700dc9847dc2bb11",
            "type": "main",
            "numbers": [
                "MainLunchHours"
            ],
            "patterns": [],
            "featurecode": false
        },
        {
            "id": "89ab5ec550f5727474f5bd5bfaa581e3",
            "name": "John Dow SmartPBX's Callflow",
            "type": "mainUserCallflow",
            "numbers": [
                "7202"
            ],
            "patterns": [],
            "featurecode": false,
            "owner_id": "d201633c77337fc469302947a56f4c44"
        },
        {
            "id": "6ec8a14cdb9db9e521ab59694c0a675e",
            "name": "ALL",
            "numbers": [
                "+74950001010",
                "+735100002020"
            ],
            "patterns": [],
            "featurecode": false
        },
        {
            "id": "63744988b3cad12ddcea42ad8f21d33b",
            "name": "all_managers_busy",
            "numbers": [
                "3",
                "all_managers_busy"
            ],
            "patterns": [],
            "featurecode": false
        },
        {
            "id": "5d6b610ac95b83bc30b19ba0aa93c057",
            "numbers": [],
            "patterns": [
                "^\\*66([0-9]{4})$"
            ],
            "featurecode": {
                "name": "eavesdrop_feature",
                "number": "66"
            }
        },
        {
            "id": "5d0a7bbf7b2a02ddda92ab4ad9235188",
            "numbers": [
                "*4"
            ],
            "patterns": [],
            "featurecode": {
                "name": "valet",
                "number": "4"
            }
        },
        {
            "id": "49ca085fe2daadee791a7c847e3e2069",
            "numbers": [],
            "patterns": [
                "^\\*0([0-9]*)$"
            ],
            "featurecode": {
                "name": "intercom",
                "number": "0"
            }
        },
        {
            "id": "4973b1f5cbd8e29f5dd0d0918a2f979c",
            "name": "Margaret Dow SmartPBX's Callflow",
            "type": "mainUserCallflow",
            "numbers": [
                "7403",
                "103",
                "+19700003875"
            ],
            "patterns": [],
            "featurecode": false,
            "owner_id": "fdb0774bf6f1c523c79a82ab4f8821bd"
        },
        {
            "id": "44d7e8a15957b18087017fd7fb544bca",
            "type": "main",
            "numbers": [
                "MainAfterHoursMenu"
            ],
            "patterns": [],
            "featurecode": false
        },
        {
            "id": "447b873bbd2a98a423334531fec32fee",
            "name": "MainConference",
            "type": "conference",
            "numbers": [
                "undefinedconf"
            ],
            "patterns": [],
            "featurecode": false
        },
        {
            "id": "43d0fde79644dea75af3bbbdd8d8087f",
            "numbers": [
                "7779"
            ],
            "patterns": [],
            "featurecode": false
        },
        {
            "id": "40c58247c8c277206b5d02d8f77e22cb",
            "numbers": [
                "+79400000001"
            ],
            "patterns": [],
            "featurecode": false
        },
        {
            "id": "3e933f5f775a28d66611cf937f2e2137",
            "name": "Irma Dow SmartPBX's Callflow",
            "type": "mainUserCallflow",
            "numbers": [
                "7410"
            ],
            "patterns": [],
            "featurecode": false,
            "owner_id": "03aee85eeb5a47c3401d86fd7f492eb3"
        },
        {
            "id": "320c2e03531adda909c9583ec34f2862",
            "name": "03 EN",
            "numbers": [
                "03"
            ],
            "patterns": [],
            "featurecode": false
        },
        {
            "id": "2c62c64888301a042258c1d2f35786fb",
            "name": "global numbers",
            "numbers": [
                "9009"
            ],
            "patterns": [],
            "featurecode": false
        },
        {
            "id": "2a69a3fbace2c4a5680825f9da8f7e8e",
            "name": "RUS_MENU",
            "numbers": [
                "2002"
            ],
            "patterns": [],
            "featurecode": false
        },
        {
            "id": "1ca30ded262c1721af91ac87e5c7a06a",
            "numbers": [
                "7777"
            ],
            "patterns": [],
            "featurecode": false
        },
        {
            "id": "1c55c08591b1325271a2839e223875ee",
            "numbers": [
                "*13"
            ],
            "patterns": [],
            "featurecode": {
                "name": "hotdesk[action=toggle]",
                "number": "13"
            }
        },
        {
            "id": "127005e49228a7f2f5ccce9594362cc8",
            "numbers": [
                "no_match"
            ],
            "patterns": [],
            "featurecode": false
        },
        {
            "id": "00a62f6212b7f8dcc3bfb012f691705c",
            "numbers": [
                "+74950000013"
            ],
            "patterns": [],
            "featurecode": false
        }
    ],
    "revision": "d3e387ba0a12087ee8055cb06995f0d3",
    "timestamp": "2019-08-17T05:43:00",
    "version": "4.2.33",
    "node": "o51LI7TmYyGMzigLbo-upw",
    "request_id": "a71f40e6093837ae0432c0e86abe0712",
    "status": "success",
    "auth_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjA4ZDQ0OTM1NDg0MTBhNDJkYjE1MWU1NTM0NjljODljIn0.eyJpc3MiOiJrYXpvbyIsImlkZW50aXR5X3NpZyI6Ikt3WFVMMHk0dWt6cWhiVWU0Y1Z2ZHNUb1Y4SVQ2NGpEMFdRc3Q4aVVucjgiLCJhY2NvdW50X2lkIjoiZmUwYWRlNDAwMDE1MzY3ZjAwNjlkNmRmYmRjYTA3MmEiLCJtZXRob2QiOiJjYl9hcGlfYXV0aCIsImV4cCI6MTU2NjAyNDEwNX0.d2CP5F7XP4QbblrCcRlfDBXxjreiOX_K6X34GPo_tzG-ct0S-GhHjxpaBAPm7YY1HmHAODaso94wEm1zzJK-tzZVrAcvJCKXD2pxR7FCUD76XL_ZJh2zXIxXbf6lxBMejKILBQvw1Oq3yPihyN1WPa61RKuPkXqbbGL1cnUesx6RqFlD6tvRxMIevV27U1y6UaYqM5JyT1u37HseCyM5XAvssMwlTFkbrdh0Oh4gs1w0SeYS63_H5IeeqLmx-QWmjmemCbh5jNAWECu1ZWxyisqhDx2GhN8SmXh1qZ2qBABhSV4O1_f-hpYorNZcr5tLbMjxBG9J9Br_tbsqvQmAxQ"
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

	resp, err := clt.CallflowsAPI.ListCallflows(ctx, "qe0ade400015367f0069d6dfbdca072a", true)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "9e1e5f9031e9e8446f54da9df47680a0", resp[0].ID, "ID's should be equal" )
	assert.ElementsMatch(t, []string{}, resp[0].Numbers, "Should be empty list")
}