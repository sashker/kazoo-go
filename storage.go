package kazooapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
)

type StorageAPIService service

type (
	Storage struct {
		ID          string                 `json:"id"`
		Attachments map[string]Attachments `json:"attachments,omitempty"`
		Connections interface{}            `json:"connections,omitempty"`
		Plan        Plan                   `json:"plan,omitempty"`
	}

	CallRecordingAttachment map[string]TypeAttachment
	MailboxMessageAttachment map[string]TypeAttachment

	Attachments struct {
		Name     string        `json:"name,omitempty"`
		Handler  string        `json:"handler"` //required
		Settings AttachmentAWS `json:"settings,omitempty"`
	}

	Plan struct {
		Account interface{}            `json:"account,omitempty"`
		Modb    Modb `json:"modb,omitempty"`
		System  interface{}            `json:"system,omitempty"`
	}

	Modb struct {
		Types map[string]TypeAttachment `json:"types,omitempty"`
	}

	TypeAttachment struct {
		Attachments TypeAttachmentHandler `json:"attachments,omitempty"`
	}

	TypeAttachmentHandler struct {
		Handler string `json:"handler,omitempty"`
	}

	TypeMailboxMessage struct {
		Attachments Attachments `json:"attachments"`
	}

	AttachmentAWS struct {
		Bucket             string      `json:"bucket"`               //required
		BucketAccessMethod string      `json:"bucket_access_method"` //required (auto vhost path)
		BucketAfterHost    bool        `json:"bucket_after_host,omitempty"`
		Host               string      `json:"host,omitempty"`
		Key                string      `json:"key"` //required
		Port               int64       `json:"port,omitempty"`
		Region             string      `json:"region"`
		Scheme             string      `json:"scheme,omitempty"`
		Secret             string      `json:"secret"`   //required
		//Settings           interface{} `json:"settings"` //required
	}
)

func (api *StorageAPIService) GetStorage(ctx context.Context, acc string) (stor *Storage, err error) {
	var response struct {
		Data Storage `json:"data"`
		ResponseEnvelope
	}

	if acc == "" {
		return nil, reportError("account id is required field")
	}

	params := Request{
		CTX:    ctx,
		Method: "GET",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/storage",
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

	stor = &response.Data

	return stor, nil
}

func (api *StorageAPIService) CreateStorage(ctx context.Context, acc string, input *Storage) (stor *Storage, err error) {
	var response struct {
		Data Storage `json:"data"`
		ResponseEnvelope
	}

	params := Request{
		CTX:    ctx,
		Method: "PUT",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/storage",
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

	stor = &response.Data

	return stor, nil
}

func (api *StorageAPIService) DeleteStorage(ctx context.Context, acc string) (stor *Storage, err error) {
	var response struct {
		Data Storage `json:"data"`
		ResponseEnvelope
	}

	if acc == "" {
		return nil, reportError("account id is required field")
	}

	params := Request{
		CTX:    ctx,
		Method: "DELETE",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/storage",
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

	stor = &response.Data

	return stor, nil
}
