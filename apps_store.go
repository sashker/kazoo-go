package kazooapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
)

type AppsStoreAPIService service

type (
	InstallAppInput struct {
		AllowedUsers string   `json:"allowed_users"`
		Users        []string `json:"users"`
	}

	InstallAppOutput struct {
		Name         string   `json:"name"`
		AllowedUsers string   `json:"allowed_users"`
		Users        []string `json:"users"`
	}
)

//InstallApp activates chosen application for given account ID
func (api *AppsStoreAPIService) InstallApp(ctx context.Context, acc, appID string, input *InstallAppInput) (output *InstallAppOutput, err error) {
	var response struct {
		Data InstallAppOutput `json:"data"`
		ResponseEnvelope
	}

	if len(appID) != 32 {
		return nil, reportError("number must be 32 symbols")
	}

	params := Request{
		CTX:    ctx,
		Method: "PUT",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/apps_store/" + appID,
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
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, reportError("Status: %v, Body: %s", resp.Status, bodyBytes)
	}

	decoder := json.NewDecoder(resp.Body)

	decErr := decoder.Decode(&response)
	if decErr != nil {
		return nil, reportError("Can't decode response: %v", decErr)
	}

	output = &response.Data

	return output, nil
}