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


func TestAccountsAPIService_ListChildren(t *testing.T) {

	ctx := context.Background()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

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

	assert.Equal(t, "2669be1c6c2d3ead16bbdd0b97aa2744", resp[0].ID, "should be equal")
}