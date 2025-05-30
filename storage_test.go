package kazooapi_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	kazooapi "github.com/sashker/kazoo-go"
	"github.com/stretchr/testify/assert"
)

func TestStorageAPIService_GetStorage(t *testing.T) {
	ctx := context.Background()

	mux := http.NewServeMux()
	mux.HandleFunc("/v2/accounts/qe0ade400015367f0069d6dfbdca072a/storage", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Server", "Cowboy")
		w.Header().Add("Content-Language", "en")
		w.Header().Add("Etag", "6a5ea7b838e7b91c5a38bc2b1050a4fa")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Vary", "accept-language, accept")
		w.Header().Add("Date", "Mon, 12 Aug 2019 05:18:39 GMT")

		body := `{
    "data": {
        "plan": {
            "modb": {
                "types": {
                    "call_recording": {
                        "attachments": {
                            "handler": "859619ec3b764362982d76b2919b602d"
                        }
                    },
                    "mailbox_message": {
                        "attachments": {
                            "handler": "859619ec3b764362982d76b2919b602d"
                        }
                    }
                }
            }
        },
        "attachments": {
            "859619ec3b764362982d76b2919b602d": {
                "settings": {
                    "secret": "iI4ytpmt5PavkxbBOgwrg45ghwegQsN56CkJw0yd",
                    "key": "IKMAJMBXMUDYQWTSGOPM",
                    "bucket": "testbucket",
                    "scheme": "https",
                    "region": "eu-central-1"
                },
                "name": "Kazoo S3",
                "handler": "s3"
            }
        },
        "id": "733027f8678d3039c02f835ec75f8e04"
    },
    "revision": "4-346061876c842dec902b776130b588f7",
    "timestamp": "2019-08-19T08:04:15",
    "version": "4.2.33",
    "node": "o51LI7TmYyGMzigLbo-upw",
    "request_id": "083f17278c2e9d6dbdd7d8ecda303520",
    "status": "success",
    "auth_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjA4ZDQ0OTM1NDg0MTBhNDJkYjE1MWU1NTM0NjljODljIn0.eyJpc3MiOiJrYXpvbyIsImlkZW50aXR5X3NpZyI6Ikt3WFVMMHk0dWt6cWhiVWU0Y1Z2ZHNUb1Y4SVQ2NGpEMFdRc3Q4aVVucjgiLCJhY2NvdW50X2lkIjoiZmUwYWRlNDAwMDE1MzY3ZjAwNjlkNmRmYmRjYTA3MmEiLCJtZXRob2QiOiJjYl9hcGlfYXV0aCIsImV4cCI6MTU2NjIwNTQzMn0.Z11jm2hI2-7OyjfowWVd_1DrZL_cgheScqzYVW5_WnSwQxpNt6zNzqOPaz51k_CIrZwJiawZYy93Z3YIZ_anqmRVf2ptPpvHGW0T1ElyzrryrWfMiEgX0hhB0grcUe5lQyer8wvMQaweSaGUABQvQpK6fh6pb5TI6rJRP0SxgJ2yQm2knjKPljT-A4_8uU78ifW-5l4GJy6fAzFkGt_3ebqFCpNHS7dsBv4G0ftorOuYQ_4VufrJPmqtOKai7Bx9DMmvzQn3UEUAeF30zesT9XGOkxP0LhRxSDMePKsmgP_uwbpzkwdYDBcZRdDNLUAqMxzPV-NB6olbij6WhOyiAw"
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

	clt, err := kazooapi.NewAPIClient(cfg)
	if err != nil {
		t.Error("Can't create the API client")
	}

	//input := &kazooapi.Account{ID: "qe0ade400015367f0069d6dfbdca072a"}

	resp, err := clt.StorageAPI.GetStorage(ctx, "qe0ade400015367f0069d6dfbdca072a")
	if err != nil {
		t.Error(err)
	}

	t.Logf("Response data %#v", resp)

	assert.Equal(t, "733027f8678d3039c02f835ec75f8e04", resp.ID, "ID's should be equal")
}

func TestStorageAPIService_CreateStorage(t *testing.T) {
	ctx := context.Background()

	mux := http.NewServeMux()
	mux.HandleFunc("/v2/accounts/qe0ade400015367f0069d6dfbdca072a/storage", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Server", "Cowboy")
		w.Header().Add("Content-Language", "en")
		w.Header().Add("Etag", "6a5ea7b838e7b91c5a38bc2b1050a4fa")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Vary", "accept-language, accept")
		w.Header().Add("Date", "Mon, 12 Aug 2019 05:18:39 GMT")

		body := `{
    "data": {
        "plan": {
            "modb": {
                "types": {
                    "call_recording": {
                        "attachments": {
                            "handler": "0f676ff8946343e797c1c46f1ffccd02"
                        }
                    },
                    "mailbox_message": {
                        "attachments": {
                            "handler": "0f676ff8946343e797c1c46f1ffccd02"
                        }
                    }
                }
            }
        },
        "attachments": {
            "0f676ff8946343e797c1c46f1ffccd02": {
                "settings": {
                    "secret": "ZZAWsHIV3hB1C9R4Ps8RDBhB0Qrl8IZMAGzzhSka",
                    "key": "MKIAJDDLLXM3XBRGOG2G",
                    "bucket": "test-bucket",
                    "scheme": "https",
                    "region": "eu-west-1"
                },
                "name": "Test",
                "handler": "s3"
            }
        },
        "id": "c1e482623df05d97074f531977866e16"
    },
    "revision": "1-78ee6fa30cdbe724cc06dc32bbfa32e1",
    "timestamp": "2019-08-19T11:45:41",
    "version": "4.2.33",
    "node": "o51LI7TmYyGMzigLbo-upw",
    "request_id": "aac51cf3d34a1be2eaf96cb8f2e4e6db",
    "status": "success",
    "auth_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjA4ZDQ0OTM1NDg0MTBhNDJkYjE1MWU1NTM0NjljODljIn0.eyJpc3MiOiJrYXpvbyIsImlkZW50aXR5X3NpZyI6Ikt3WFVMMHk0dWt6cWhiVWU0Y1Z2ZHNUb1Y4SVQ2NGpEMFdRc3Q4aVVucjgiLCJhY2NvdW50X2lkIjoiZmUwYWRlNDAwMDE1MzY3ZjAwNjlkNmRmYmRjYTA3MmEiLCJtZXRob2QiOiJjYl9hcGlfYXV0aCIsImV4cCI6MTU2NjIxODY0Mn0.EZ0JV3VEX9_HeEzvumpIK0UUbK4itVYWfSn6QLCxLNwaU8dYLtuTcw7P9u5UQCyyM7rTqBWv4jCuzaEsaFyG0hMx1awnMEo7Tlwr_Cp3ruxMof5zFI4cxYybXbTnR1H30VP96w-Q-OSNB2we5KUj7nF0hTRqZe_DCJLYLaR5vpx5XVZBdgN2Ihjd8nTHxYZeJk1OvYYqKT_Mq6_xueSOT_r7xbdqiBhGeNwgkNClGdBva_dHVFiyf8v8R0Zrp-OUplAt6u0WJZftDSMPElz32h77IhVPyJb3TwFAcfg_tNZ4Zc8W_O_Jw39j063qdCewjacOoG-Qz6MyQwckOofDPg"
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

	clt, err := kazooapi.NewAPIClient(cfg)
	if err != nil {
		t.Error("Can't create the API client")
	}

	input := &kazooapi.Storage{
		Attachments: map[string]kazooapi.Attachments{
			"0f676ff8946343e797c1c46f1ffccd02": kazooapi.Attachments{
				Name:    "Test",
				Handler: "s3",
				Settings: kazooapi.AttachmentAWS{
					Secret: "ZZAWsHIV3hB1C9R4Ps8RDBhB0Qrl8IZMAGzzhSka",
					Key:    "MKIAJDDLLXM3XBRGOG2G",
					Bucket: "test-bucket",
					Scheme: "https",
					Region: "eu-west-1",
				},
			},
		},
		Plan: kazooapi.Plan{
			Account: nil,
			Modb: kazooapi.Modb{
				Types: map[string]kazooapi.TypeAttachment{
					"call_recording": kazooapi.TypeAttachment{
						Attachments: kazooapi.TypeAttachmentHandler{
							Handler: "0f676ff8946343e797c1c46f1ffccd02",
						},
					},
					"mailbox_message": kazooapi.TypeAttachment{
						Attachments: kazooapi.TypeAttachmentHandler{
							Handler: "0f676ff8946343e797c1c46f1ffccd02",
						},
					},
				},
			},
			System: nil,
		},
	}

	resp, err := clt.StorageAPI.CreateStorage(ctx, "qe0ade400015367f0069d6dfbdca072a", input)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Response data %#v", resp)

	assert.Equal(t, "c1e482623df05d97074f531977866e16", resp.ID, "ID's should be equal")
}
