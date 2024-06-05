package kazooapi_test

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	kazooapi "gitlab.com/bmitelecom/kazoo-go"
)

func TestAccountsAPIService_ListChildren(t *testing.T) {

	ctx := context.Background()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Server", "Cowboy")
		w.Header().Add("Content-Language", "en")
		w.Header().Add("Etag", "6a5ea7b838e7b91c5a38bc2b1050a4fa")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Vary", "accept-language, accept")
		w.Header().Add("Date", "Mon, 12 Aug 2019 05:18:39 GMT")

		body := `{
    "next_start_key": "g20AAAADMTUy",
    "page_size": 50,
    "start_key": "g20AAAAA",
    "data": [
        {
            "id": "2669be1c6c2d3ead16bbdd0b97aa2744",
            "name": "1002",
            "realm": "1002.pbx.example.com",
            "descendants_count": 0
        },
        {
            "id": "847c95e6cce8114beebbc69fc962aa32",
            "name": "1102",
            "realm": "1102.pbx.example.com",
            "descendants_count": 0
        },
        {
            "id": "8da6414d38009540d57993abcc0aabaa",
            "name": "1157",
            "realm": "1157.pbx.example.com",
            "descendants_count": 0
        },
        {
            "id": "76a37a8e684b23348660119cf6e06a65",
            "name": "12345",
            "realm": "12345.pbx.example.com",
            "descendants_count": 0
        },
        {
            "id": "d5972cbb318a8a7fcb16a5ea3c63040d",
            "name": "1239",
            "realm": "1239.pbx.example.com",
            "descendants_count": 0
        },
        {
            "id": "b339f45e293d3615e31ec49a106df21c",
            "name": "1516",
            "realm": "1516.pbx.example.com",
            "descendants_count": 0
        }
    ],
    "revision": "6a5ea7b838e7b91c5a38bc2b1050a4fa",
    "timestamp": "2019-08-12T05:18:39",
    "version": "4.2.33",
    "node": "o51LI7TmYyGMzigLbo-upw",
    "request_id": "3ae535b86688c03169207e2185d74aab",
    "status": "success",
    "auth_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjA4ZDQ0OTM1NDg0MTBhNDJkYjE1MWU1NTM0NjljODljIn0.eyJpc3MiOiJrYXpvbyIsImlkZW50aXR5X3NpZyI6Ikt3WFVMMHk0dWt6cWhiVWU0Y1Z2ZHNUb1Y4SVQ2NGpEMFdRc3Q4aVVucjgiLCJhY2NvdW50X2lkIjoiZmUwYWRlNDAwMDE1MzY3ZjAwNjlkNmRmYmRjYTA3MmEiLCJtZXRob2QiOiJjYl9hcGlfYXV0aCIsImV4cCI6MTU2NTU5MDY5N30.f0K9R9NjAc27kbxv5Hob543pZlx064C_RVWUc4oIcD9fNH5IWkkDD_l6RUnm9VWHzPftN71IKvIzUrPeHH4db5xg8IZipEunvIzsY6dIxW_s74tvf06UXPbugIv1za0w8tC1SrEQGXNHOW-yFGjiu67NK36XcvsCgPZNoRICND1GoJDt33xGk4X2KzD21vltZRUUdViwPUvrhjQnK2sM6lF0CzBbOZaEfCiTuLgwL8CTliT5aaTEW84XvuY4r3_rjAb7PAOe9524tccQWrYjz5hiRfGUw93RZQ4SUzyuER4GbZm1kNk-7FG1RR5kz5N8zhnsoVe10EuyJ2--V191Zw"
}`
		io.WriteString(w, body)

	})

	srv := httptest.NewServer(mux)
	defer srv.Close()

	cfg := kazooapi.NewConfiguration()
	cfg.APIKey = "e0a582bad3fb7fe3897ebf70cc0f542bbdc9a17895764266f094b953254d3d84"
	cfg.BasePath = srv.URL + "/accounts/qe0ade400015367f0069d6dfbdca072a/children"

	clt, err := kazooapi.NewAPIClient(cfg)
	if err != nil {
		t.Error("Can't create the API client")
	}

	//input := &kazooapi.Account{ID: "qe0ade400015367f0069d6dfbdca072a"}

	resp, err := clt.AccountsAPI.ListChildren(ctx, "qe0ade400015367f0069d6dfbdca072a", true)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)

	//assert.Equal(t, "2669be1c6c2d3ead16bbdd0b97aa2744", resp[0].ID, "should be equal")
}

func TestAccountsAPIService_ChangeAccount(t *testing.T) {

	ctx := context.Background()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, `{"data":{"call_recording":{"account":{"any":{"any":{"enabled":true,"format":"mp3"}}}}}}`, string(b), "should be equal")

		w.Header().Add("Server", "Cowboy")
		w.Header().Add("Content-Language", "en")
		w.Header().Add("Etag", "6a5ea7b838e7b91c5a38bc2b1050a4fa")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Vary", "accept-language, accept")
		w.Header().Add("Date", "Mon, 12 Aug 2019 05:18:39 GMT")

		body := `{
    "data": {
        "ui_restrictions": {
            "myaccount": {
                "user": {
                    "show_tab": true
                },
                "twoway": {
                    "show_tab": true
                },
                "transactions": {
                    "show_tab": true
                },
                "service_plan": {
                    "show_tab": true
                },
                "outbound": {
                    "show_tab": true
                },
                "inbound": {
                    "show_tab": true
                },
                "errorTracker": {
                    "show_tab": true
                },
                "billing": {
                    "show_tab": true
                },
                "balance": {
                    "show_tab": true,
                    "show_minutes": true,
                    "show_credit": true
                },
                "account": {
                    "show_tab": true
                }
            }
        },
        "ui_metadata": {
            "version": "4.3.0",
            "ui": "monster-ui",
            "origin": "callflows"
        },
        "timezone": "Europe/Moscow",
        "ringtones": {},
        "reseller_id": "fe0ade400015367f0069d6dfbdca072a",
        "realm": "test.pbx.example.com",
        "preflow": {},
        "notifications": {
            "low_balance": {
                "sent_low_balance": true,
                "last_notification": 63747348527
            },
            "first_occurrence": {
                "sent_initial_registration": true,
                "sent_initial_call": false
            }
        },
        "name": "TEST",
        "music_on_hold": {},
        "is_reseller": false,
        "created": 63747343047,
        "caller_id": {
            "internal": {
                "number": "",
                "name": ""
            },
            "external": {
                "number": "+12345678999"
            },
            "emergency": {
                "number": "+12345678999",
                "name": ""
            }
        },
        "call_restriction": {
            "unknown": {
                "action": "inherit"
            },
            "toll_ru": {
                "action": "inherit"
            },
            "service": {
                "action": "inherit"
            },
            "russia_mobile": {
                "action": "inherit"
            },
            "russia_land": {
                "action": "inherit"
            },
            "international": {
                "action": "inherit"
            },
            "SNG": {
                "action": "inherit"
            }
        },
        "call_recording": {
            "account": {
                "any": {
                    "any": {
                        "format": "mp3",
                        "enabled": true
                    }
                }
            }
        },
        "blacklists": [],
        "id": "9235595b763c59359bde2a277973f1e8",
        "wnm_allow_additions": false,
        "superduper_admin": false,
        "enabled": true,
        "billing_mode": "limits_only"
    },
    "revision": "6-89f48cf5afb7ec87ea444d18462d3def",
    "timestamp": "2020-01-29T08:06:38",
    "version": "4.2.33",
    "node": "o51LI7TmYyGMzigLbo-upw",
    "request_id": "7b75baa919b7c9ccad5325e0b2c645e9",
    "status": "success",
    "auth_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjA4ZDQ0OTM1NDg0MTBhNDJkYjE1MWU1NTM0NjljODljIn0.eyJpc3MiOiJrYXpvbyIsImlkZW50aXR5X3NpZyI6Ikt3WFVMMHk0dWt6cWhiVWU0Y1Z2ZHNUb1Y4SVQ2NGpEMFdRc3Q4aVVucjgiLCJhY2NvdW50X2lkIjoiZmUwYWRlNDAwMDE1MzY3ZjAwNjlkNmRmYmRjYTA3MmEiLCJtZXRob2QiOiJjYl9hcGlfYXV0aCIsImV4cCI6MTU4MDI4ODc5NH0.ALazRrD-S-inTag8JxE30hE7qtO-Y3-JKXu-0aWl1JPU6eSoAvYyMgSXOyytRIx3hCBag-lL3qdfk7Pchy-5hr7RO6bcRpOcyT7ArqAlLwGBHRV30Zp9cSRNbwmaB-zpY0zdEZ-mIsSKR9HLYibQZalyIAvxaGVdu-CVKN_jQd0kGtxGKrvJhVdXueulARdl40mhwK4mLmzsykwLVyDa-HwkvlrjZ9ruSGJhhhhl_lx8Efa9IlYqsACy3tKFtZd70YBWG4Z4h6e3sGatYzza4TOFw-0DPwRmEDLtOokO6MK0mXsFX7oAQZhQsBX-sqd1qch21ClehteC5KU1_6tUjQ"
}`
		io.WriteString(w, body)

	})
	mux.HandleFunc("/v2/api_auth", func(w http.ResponseWriter, r *http.Request) {

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
	//cfg.BasePath = srv.URL + "/accounts/qe0ade400015367f0069d6dfbdca072a"

	clt, err := kazooapi.NewAPIClient(cfg)
	if err != nil {
		t.Error("Can't create the API client")
	}

	//input := &kazooapi.Account{ID: "qe0ade400015367f0069d6dfbdca072a"}

	input := map[string]interface{}{
		"call_recording": map[string]interface{}{
			"account": map[string]interface{}{
				"any": map[string]interface{}{
					"any": map[string]interface{}{
						"enabled": true,
						"format":  "mp3",
					},
				},
			},
		},
	}

	resp, err := clt.AccountsAPI.ChangeAccount(ctx, "9235595b763c59359bde2a277973f1e8", input)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "9235595b763c59359bde2a277973f1e8", resp.ID, "should be equal")
}

func Test_EnableCallRecording_Success(t *testing.T) {
	ctx := context.Background()
	cfg := kazooapi.NewConfiguration()
	cfg.APIKey = "e0a582bad3fb7fe3897ebf70cc0f542bbdc9a17895764266f094b953254d3d84"
	cfg.BasePath = "http://kazoo2.hz.sip3.net:8000" + "/v2"
	//cfg.BasePath = srv.URL + "/accounts/qe0ade400015367f0069d6dfbdca072a"

	clt, err := kazooapi.NewAPIClient(cfg)
	if err != nil {
		t.Error("Can't create the API client")
	}

	//input := &kazooapi.Account{ID: "qe0ade400015367f0069d6dfbdca072a"}

	input := map[string]interface{}{
		"call_recording": map[string]interface{}{
			"account": map[string]interface{}{
				"any": map[string]interface{}{
					"any": map[string]interface{}{
						"enabled": true,
						"format":  "mp3",
					},
				},
			},
		},
	}

	resp, err := clt.AccountsAPI.ChangeAccount(ctx, "4dee5c1bef3ace50911c9917c50c9f80", input)
	assert.NoError(t, err)

	assert.Equal(t, "9235595b763c59359bde2a277973f1e8", resp.ID, "should be equal")
}
