package kazooapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
)

type AccountsAPIService service

type (
	Account struct {
		ID                string      `json:"id,omitempty"`
		Enabled           bool        `json:"enabled,omitempty"`
		Name              string      `json:"name"`
		BillingMode       string      `json:"billing_mode,omitempty"`
		CallRestriction   interface{} `json:"call_restriction,omitempty"`
		CallerID          interface{} `json:"caller_id,omitempty"`
		Created           int64       `json:"created,omitempty"`
		DialPlan          interface{} `json:"dial_plan,omitempty"`
		IsReseller        bool        `json:"is_reseller,omitempty"`
		Language          string      `json:"language,omitempty"`
		MusicOnHold       interface{} `json:"music_on_hold,omitempty"`
		Preflow           interface{} `json:"preflow,omitempty"`
		Realm             string      `json:"realm,omitempty"`
		ResellerID        string      `json:"reseller_id,omitempty"`
		Ringtones         interface{} `json:"ringtones,omitempty"`
		SuperduperAdmin   bool        `json:"superduper_admin,omitempty"`
		Timezone          string      `json:"timezone,omitempty"`
		WnmAllowAdditions bool        `json:"wnm_allow_additions,omitempty"`
		Org               string      `json:"org,omitempty"`
	}

	Child struct {
		ID               string `json:"id"`
		Name             string `json:"name"`
		Realm            string `json:"realm"`
		DescendantsCount int    `json:"descendants_count"`
	}

	Descendant struct {
		ID    string   `json:"id"`
		Name  string   `json:"name"`
		Realm string   `json:"realm"`
		Tree  []string `json:"tree"`
	}
)

func (api *AccountsAPIService) GetAccount(ctx context.Context, id string) (acc *Account, err error) {
	var response struct {
		Data Account `json:"data"`
		ResponseEnvelope
	}

	if id == "" {
		return nil, reportError("account id is required field")
	}

	params := Request{
		CTX:    ctx,
		Method: "GET",
		Path:   api.client.cfg.BasePath + "/accounts/" + id,
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

	acc = &response.Data

	return acc, nil
}

func (api *AccountsAPIService) CreateAccount(ctx context.Context, input *Account) (acc *Account, err error) {
	var (
		response struct {
			Data Account `json:"data"`
			ResponseEnvelope
		}
	)

	if input.Name == "" {
		return nil, reportError("account name is required field")
	}

	/*if input.Realm == "" {
		return nil, reportError("realm is required field")
	}*/

	params := Request{
		CTX:    ctx,
		Method: "PUT",
		Path:   api.client.cfg.BasePath + "/accounts",
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
	if err != nil {
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

	acc = &response.Data

	return acc, nil
}

//ChangeAccount enables to PATCH an existing account document
func (api *AccountsAPIService) ChangeAccount(ctx context.Context, id string, input map[string]interface{}) (acc *Account, err error) {
	var (
		response struct {
			Data Account `json:"data"`
			ResponseEnvelope
		}
	)

	if id == "" || len(id) != 32 {
		return nil, reportError("specify correct account id")
	}

	/*if input.Realm == "" {
		return nil, reportError("realm is required field")
	}*/

	params := Request{
		CTX:    ctx,
		Method: "PATCH",
		Path:   api.client.cfg.BasePath + "/accounts/" + id,
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

	acc = &response.Data

	return acc, nil
}

func (api *AccountsAPIService) ListChildren(ctx context.Context, acc string, disablePagination bool) (chldrn []Child, err error) {
	var response struct {
		Data []Child `json:"data"`
		ResponseEnvelope
	}

	params := Request{
		CTX:    ctx,
		Method: "GET",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/children",
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

	chldrn = response.Data

	return chldrn, nil
}

func (api *AccountsAPIService) ListDescendants(ctx context.Context, acc string, disablePagination bool) (chldrn []Descendant, err error) {
	var response struct {
		Data []Descendant `json:"data"`
		ResponseEnvelope
	}

	params := Request{
		CTX:    ctx,
		Method: "GET",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/descendants",
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

	chldrn = response.Data

	return chldrn, nil
}
