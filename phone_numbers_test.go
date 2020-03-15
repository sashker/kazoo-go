package kazooapi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	kazooapi "gitlab.com/bmitelecom/kazoo-go"
)

func TestPhoneNumbersService_DeletePhoneNumber(t *testing.T) {
	ctx := context.Background()

	srv := MockKazooServer(t, "")

	cfg := kazooapi.NewConfiguration()
	cfg.APIKey = "e0a582bad3fb7fe3897ebf70cc0f542bbdc9a17895764266f094b953254d3d84"
	cfg.BasePath = srv.URL + "/v2"
	cfg.HTTPClient = srv.Client()

	clt, err := kazooapi.NewAPIClient(cfg)
	if err != nil {
		t.Error("Can't create the API client")
	}

	//input := &kazooapi.Account{ID: "qe0ade400015367f0069d6dfbdca072a"}

	resp, err := clt.PhoneNumbersAPI.DeletePhoneNumber(ctx, "qe0ade400015367f0069d6dfbdca072a", "+74955555555")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "+74955555555", resp.ID, "ID's should be equal")
	//assert.ElementsMatch(t, []string{}, resp[0].Numbers, "Should be empty list")
}
