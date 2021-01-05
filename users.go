package kazooapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
)

type UsersAPIService service

type (
	User struct {
		ID         string   `json:"id,omitempty"`
		Username   string   `json:"username,omitempty"`
		Password   string   `json:"password,omitempty"`
		FirstName  string   `json:"first_name,omitempty"`
		LastName   string   `json:"last_name,omitempty"`
		Title      string   `json:"title,omitempty"`
		Email      string   `json:"email,omitempty"`
		PresenceID string   `json:"presence_id,omitempty"`
		Enabled    bool     `json:"enabled,omitempty"`
		Features   []string `json:"features,omitempty"`
		Verified   bool     `json:"verified,omitempty"`
		PrivLevel  string   `json:"priv_level,omitempty"`
		Timezone   string   `json:"timezone,omitempty"`
	}
)

func (api *UsersAPIService) CreateUser(ctx context.Context, acc string, input *User) (usr *User, err error) {
	var response struct {
		Data User `json:"data"`
		ResponseEnvelope
	}

	if acc == "" {
		return nil, reportError("account id is required field")
	}

	if input.FirstName == "" {
		return nil, reportError("first name is required field")
	}

	if input.LastName == "" {
		return nil, reportError("last name is required field")
	}

	params := Request{
		CTX:    ctx,
		Method: "PUT",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/users",
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

	usr = &response.Data

	return usr, nil
}

func (api *UsersAPIService) DeleteUser(ctx context.Context, acc, id string) (usr *User, err error) {
	var response struct {
		Data User `json:"data"`
		ResponseEnvelope
	}

	if acc == "" {
		return nil, reportError("account id is required field")
	}

	params := Request{
		CTX:    ctx,
		Method: "DELETE",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/users/" + id,
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

	usr = &response.Data

	return usr, nil
}

func (api *UsersAPIService) ListUsers(ctx context.Context, acc string, disablePagination bool) (users []User, err error) {
	var response struct {
		Data []User `json:"data"`
		ResponseEnvelope
	}

	params := Request{
		CTX:    ctx,
		Method: "GET",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/users",
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

	decoder := json.NewDecoder(resp.Body)

	decErr := decoder.Decode(&response)
	if decErr != nil {
		return nil, reportError("Can't decode response: %v", decErr)
	}

	users = response.Data

	return users, nil
}
