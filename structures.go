package kazooapi

type (
	CallForward struct {
		Enabled          bool   `json:"enabled"`
		Number           string `json:"number"`
		DirectCallsOnly  bool   `json:"direct_calls_only"`
		Failover         bool   `json:"failover"`
		IgnoreEarlyMedia bool   `json:"ignore_early_media"`
		KeepCallerID     bool   `json:"keep_caller_id"`
		RequireKeyPress  bool   `json:"require_key_press"`
		Substitute       bool   `json:"substitute"`
	}

)
