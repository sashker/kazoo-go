package kazooapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
)

type ClicktocallAPIService service

type (
	Clicktocall struct {
		ID                     string            `json:"id"`
		Name                   string            `json:"name"`
		AuthRequired           bool              `json:"auth_required,omitempty"`
		BypassMedia            bool              `json:"bypass_media,omitempty"`
		CallerIDNumber         string            `json:"caller_id_number,omitempty"`
		CustomApplicationVars  map[string]string `json:"custom_application_vars,omitempty"`
		CustomSIPHeaders       map[string]string `json:"custom_sip_headers,omitempty"`
		DialFirst              string            `json:"dial_first,omitempty"` //extension|contact
		Extension              string            `json:"extension"`
		Media                  interface{}       `json:"media,omitempty"` //TODO
		OutboundCalleeIDName   string            `json:"outbound_callee_id_name,omitempty"`
		OutboundCalleeIDNumber string            `json:"outbound_callee_id_number,omitempty"`
		PresenceID             string            `json:"presence_id,omitempty"`
		Ringback               string            `json:"ringback,omitempty"`
		Throttle               int64             `json:"throttle,omitempty"`
		Timeout                int64             `json:"timeout,omitempty"`
		Whitelist              []string          `json:"whitelist,omitempty"`
	}

	ClicktocallExecuteResponse struct {
		ApplicationData struct {
			Route string `json:"route"`
		} `json:"application_data"`
		ApplicationName       string `json:"application_name"`
		ContinueOnFail        bool   `json:"continue_on_fail"`
		CustomApplicationVars struct {
			Contact string `json:"contact"`
		} `json:"custom_application_vars"`
		CustomChannelVars struct {
			AccountID          string `json:"account_id"`
			AuthorizingID      string `json:"authorizing_id"`
			AuthorizingType    string `json:"authorizing_type"`
			AutoAnswerLoopback bool   `json:"auto_answer_loopback"`
			FromURI            string `json:"from_uri"`
			InheritCodec       bool   `json:"inherit_codec"`
			LoopbackRequestURI string `json:"loopback_request_uri"`
			RequestURI         string `json:"request_uri"`
			RetainCid          bool   `json:"retain_cid"`
		} `json:"custom_channel_vars"`
		DialEndpointMethod string `json:"dial_endpoint_method"`
		Endpoints          []struct {
			InviteFormat string `json:"invite_format"`
			Route        string `json:"route"`
			ToDid        string `json:"to_did"`
			ToRealm      string `json:"to_realm"`
		} `json:"endpoints"`
		ExportCustomChannelVars []string `json:"export_custom_channel_vars"`
		IgnoreEarlyMedia        bool     `json:"ignore_early_media"`
		LoopbackBowout          string   `json:"loopback_bowout"`
		OutboundCallID          string   `json:"outbound_call_id"`
		OutboundCalleeIDName    string   `json:"outbound_callee_id_name"`
		OutboundCalleeIDNumber  string   `json:"outbound_callee_id_number"`
		OutboundCallerIDName    string   `json:"outbound_caller_id_name"`
		OutboundCallerIDNumber  string   `json:"outbound_caller_id_number"`
		SimplifyLoopback        string   `json:"simplify_loopback"`
		StartControlProcess     string   `json:"start_control_process"`
		Timeout                 int      `json:"timeout"`
	}
)

//GetClicktocall fetches parameters of selected clicktocall
func (api *Clicktocall) GetClicktocall(ctx context.Context, acc, id string) (c2c *Clicktocall, err error) {
	var response struct {
		Data Clicktocall `json:"data"`
		ResponseEnvelope
	}

	if id == "" {
		return nil, reportError("clicktocall id is required field")
	}

	params := Request{
		CTX:    ctx,
		Method: "GET",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/clicktocall/" + id,
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

	err = readBody(resp, &response)
	if err != nil {
		return nil, reportError("Can't decode response: %v", err)
	}

	if resp.StatusCode >= 300 {
		return nil, prepareError(resp)
	}

	c2c = &response.Data

	return c2c, nil
}

//CreateClicktocall creates a new clicktocall with given parameters
func (api *ClicktocallAPIService) CreateClicktocall(ctx context.Context, acc string, input *Clicktocall) (c2c *Clicktocall, err error) {
	var (
		response struct {
			Data Clicktocall `json:"data"`
			ResponseEnvelope
		}
	)

	if input.Name == "" {
		return nil, reportError("Clicktocall name is required field")
	}

	if input.Extension == "" {
		return nil, reportError("Extension is required field")
	}

	/*if input.Realm == "" {
		return nil, reportError("realm is required field")
	}*/

	params := Request{
		CTX:    ctx,
		Method: "PUT",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/clicktocall",
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
		return nil, reportError("Can't prepare a request %s", err)
	}

	resp, err := api.client.callAPI(ctx, req)
	if err != nil || resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return nil, prepareError(resp)
	}

	err = readBody(resp, &response)
	if err != nil {
		return nil, reportError("Can't decode response: %v", err)
	}

	c2c = &response.Data

	return c2c, nil
}

//ListClick2Calls lists all clicktocall endpoints
func (api *ClicktocallAPIService) ListClick2Calls(ctx context.Context, acc string, disablePagination bool) (result []Clicktocall, err error) {
	var response struct {
		Data []Clicktocall `json:"data"`
		ResponseEnvelope
	}

	params := Request{
		CTX:    ctx,
		Method: "GET",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/clicktocall",
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
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, reportError("Status: %v, Body: %s", resp.Status, bodyBytes)
	}

	err = readBody(resp, &response)
	if err != nil {
		return nil, reportError("Can't decode response: %v", err)
	}

	result = response.Data

	return result, nil
}

//ExecuteClicktocall executes non-blocking version of clicktocall
func (api *Clicktocall) ExecuteClicktocall(ctx context.Context, acc, id, contact string) (cer *ClicktocallExecuteResponse, err error) {
	var response struct {
		Data ClicktocallExecuteResponse `json:"data"`
		ResponseEnvelope
	}

	if id == "" {
		return nil, reportError("clicktocall id is required field")
	}

	params := Request{
		CTX:    ctx,
		Method: "GET",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/clicktocall/" + id + "/connect?contact=" + contact,
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

	err = readBody(resp, &response)
	if err != nil {
		return nil, reportError("Can't decode response: %v", err)
	}

	if resp.StatusCode >= 300 {
		return nil, prepareError(resp)
	}

	cer = &response.Data

	return cer, nil
}

//DeleteClicktocall deletes selected clicktocall
func (api *Clicktocall) DeleteClicktocall(ctx context.Context, acc, id string) (c2c *Clicktocall, err error) {
	var response struct {
		Data Clicktocall `json:"data"`
		ResponseEnvelope
	}

	if id == "" {
		return nil, reportError("clicktocall id is required field")
	}

	params := Request{
		CTX:    ctx,
		Method: "DELETE",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/clicktocall/" + id,
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

	err = readBody(resp, &response)
	if err != nil {
		return nil, reportError("Can't decode response: %v", err)
	}

	if resp.StatusCode >= 300 {
		return nil, prepareError(resp)
	}

	c2c = &response.Data

	return c2c, nil
}

//ChangeClicktocall changes clicktocall's parameters
func (api *ClicktocallAPIService) ChangeClicktocall(ctx context.Context, acc,id string, input *Clicktocall) (c2c *Clicktocall, err error) {
	var (
		response struct {
			Data Clicktocall `json:"data"`
			ResponseEnvelope
		}
	)

	if input.Name == "" {
		return nil, reportError("Clicktocall name is required field")
	}

	if input.Extension == "" {
		return nil, reportError("Extension is required field")
	}

	/*if input.Realm == "" {
		return nil, reportError("realm is required field")
	}*/

	params := Request{
		CTX:    ctx,
		Method: "POST",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/clicktocall/" + id,
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
		return nil, reportError("Can't prepare a request %s", err)
	}

	resp, err := api.client.callAPI(ctx, req)
	if err != nil || resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return nil, prepareError(resp)
	}

	err = readBody(resp, &response)
	if err != nil {
		return nil, reportError("Can't decode response: %v", err)
	}

	c2c = &response.Data

	return c2c, nil
}

