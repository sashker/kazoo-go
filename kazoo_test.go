package kazooapi_test

import (
	"github.com/stretchr/testify/assert"
	kazooapi "gitlab.com/bmitelecom/kazoo-go"
	"testing"
)

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
