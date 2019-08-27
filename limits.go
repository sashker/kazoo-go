package kazooapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
)

type LimitsAPIService service

type (
	Limits struct {
		AllowPrepay            bool     `json:"allow_prepay,omitempty"`
		AuthzResourceTypes     []string `json:"authz_resource_types,omitempty"`
		BurstTrunks            int64    `json:"burst_trunks,omitempty"`
		Calls                  int64    `json:"calls,omitempty"`
		InboundTrunks          int64    `json:"inbound_trunks,omitempty"`
		OutboundTrunks         int64    `json:"outbound_trunks,omitempty"`
		ResourceConsumingCalls int64    `json:"resource_consuming_calls,omitempty"`
		TwowayTrunks           int64    `json:"twoway_trunks,omitempty"`
	}
)

func (api *LimitsAPIService) GetLimits(ctx context.Context, acc string) (limits *Limits, err error) {
	var response struct {
		Data Limits `json:"data"`
		ResponseEnvelope
	}

	if acc == "" {
		return nil, reportError("account id is required field")
	}

	params := Request{
		CTX:    ctx,
		Method: "GET",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/limits",
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
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, reportError("Status: %v, Body: %s", resp.Status, bodyBytes)
	}

	decoder := json.NewDecoder(resp.Body)

	decErr := decoder.Decode(&response)
	if decErr != nil {
		return nil, reportError("Can't decode response: %v", decErr)
	}

	limits = &response.Data

	return limits, nil
}

func (api *LimitsAPIService) UpdateLimits(ctx context.Context, acc string, input *Limits) (limits *Limits, err error) {
	var response struct {
		Data Limits `json:"data"`
		ResponseEnvelope
	}

	params := Request{
		CTX:    ctx,
		Method: "POST",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/limits",
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
		//bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, reportError("error status code: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)

	decErr := decoder.Decode(&response)
	if decErr != nil {
		return nil, reportError("Can't decode response: %v", decErr)
	}

	limits = &response.Data

	return limits, nil
}