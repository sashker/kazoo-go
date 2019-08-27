package kazooapi

import (
	"context"
	"encoding/json"
)

type CallflowsAPIService service

type (
	Callflow struct {
		ID          string         `json:"id,omitempty"`
		Name        string         `json:"name,omitempty"`
		Type        string         `json:"type,omitempty"`
		Flags       []string       `json:"flags,omitempty"`
		Flow        CallflowAction `json:"flow"`
		Numbers     []string       `json:"numbers,omitempty"`
		Patterns    []string       `json:"patterns,omitempty"`
		FeatureCode interface{}    `json:"featurecode,omitempty"`
	}

	CallflowAction struct {
		Module   string      `json:"module"`
		Children interface{} `json:"children"`
		Data     interface{} `json:"data"`
	}

	FeatureCode struct {
		Number string `json:"number,omitempty"`
		Name   string `json:"name,omitempty"`
	}

	Metaflow struct {
		BindingDigit string      `json:"binding_digit,omitempty"`
		DigitTimeout int         `json:"digit_timeout,omitempty"`
		ListenOn     string      `json:"listen_on,omitempty"`
		Numbers      []string    `json:"numbers,omitempty"`
		Patterns     []string    `json:"patterns,omitempty"`
		FeatureCode  FeatureCode `json:"featurecode,omitempty"`
	}

	MetaflowAction struct {
		CallflowAction
	}

	Resources struct {
		BypassE164        bool     `json:"bypass_e164,omitempty"`
		CallerIDType      string   `json:"caller_id_type,omitempty"` //external - is default value
		DoNotNormalize    bool     `json:"do_not_normalize,omitempty"`
		DynamicFlags      []string `json:"dynamic_flags,omitempty"`
		EmitAccountID     bool     `json:"emit_account_id,omitempty"`
		FormatFromURI     bool     `json:"format_from_uri,omitempty"`
		FromURIRealm      string   `json:"from_uri_realm,omitempty"`
		HuntAccountID     string   `json:"hunt_account_id,omitempty"`
		IgnoreEarlyMedia  bool     `json:"ignore_early_media,omitempty"`
		OutboundFlags     []string `json:"outbound_flags,omitempty"`
		Ringback          string   `json:"ringback,omitempty"`
		Timeout           int64    `json:"timeout,omitempty"`
		ToDID             string   `json:"to_did,omitempty"`
		UseLocalResources bool     `json:"use_local_resources,omitempty"`
	}

	UserModule struct {
		ID                 string      `json:"id,omitempty"`
		SkipModule         bool        `json:"skip_module,omitempty"`
		CanCallSelf        bool        `json:"can_call_self,omitempty"`
		CanTextSelf        bool        `json:"can_text_self,omitempty"`
		CustomSIPHeaders   interface{} `json:"custom_sip_headers,omitempty"`
		Delay              int64       `json:"delay,omitempty"`
		Timeout            int64       `json:"timeout,omitempty"`
		FailOnSingleReject bool        `json:"fail_on_single_reject,omitempty"`
		StaticInvite       string      `json:"static_invite,omitempty"`
		Strategy           string      `json:"strategy,omitempty"`
		SupressCLID        bool        `json:"suppress_clid,omitempty"`
	}
)

func (api *CallflowsAPIService) CreateCallflow(ctx context.Context, acc string, input *Callflow) (cf *Callflow, err error) {
	var response struct {
		Data Callflow `json:"data"`
		ResponseEnvelope
	}

	params := Request{
		CTX:    ctx,
		Method: "PUT",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/callflows",
	}

	reqBody := RequestEnvelope{
		Data: input,
	}

	jsonString, err := json.Marshal(reqBody)
	if err != nil {
		return nil, reportError("can't marshall body for request")
	}

	body, err := setBody(jsonString, "json")
	if err != nil {
		return nil, reportError("can't prepare body for the request")
	}

	params.PostBody = body

	req, err := api.client.prepareRequest(&params)
	if err != nil {
		return nil, reportError("can't prepare a request %s", err)
	}

	resp, err := api.client.callAPI(ctx, req)
	if err != nil || resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return nil, prepareError(resp)
	}

	decoder := json.NewDecoder(resp.Body)

	decErr := decoder.Decode(&response)
	if decErr != nil {
		return nil, reportError("Can't decode response: %v", decErr)
	}

	cf = &response.Data

	return cf, nil
}

//ListCallflows shows callflows belong to a given account
func (api *CallflowsAPIService) ListCallflows(ctx context.Context, acc string, disablePagination bool) (cfs []Callflow, err error) {
	var response struct {
		Data []Callflow `json:"data"`
		ResponseEnvelope
	}

	params := Request{
		CTX:    ctx,
		Method: "GET",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/callflows",
	}

	if disablePagination {
		params.QueryParams = map[string][]string{"paginate": []string{"false"}}
	}

	req, err := api.client.prepareRequest(&params)
	if err != nil {
		return nil, reportError("Can't prepare a request %s", err)
	}

	resp, err := api.client.callAPI(ctx, req)
	if err != nil || resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		//bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, reportError("status code: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)

	decErr := decoder.Decode(&response)
	if decErr != nil {
		return nil, reportError("Can't decode response: %v", decErr)
	}

	cfs = response.Data

	return cfs, nil
}
