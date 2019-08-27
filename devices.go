package kazooapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
)

type DevicesAPIService service

type (
	Device struct {
		ID                              string      `json:"id,omitempty"`
		Name                            string      `json:"name"`
		OwnerID                         string      `json:"owner_id"`
		DeviceType                      string      `json:"device_type,omitempty"`
		Enabled                         bool        `json:"enabled,omitempty"`
		ExcludeFromQueues               bool        `json:"exclude_from_queues,omitempty"`
		Flags                           []string    `json:"flags,omitempty"`
		MACAddress                      string      `json:"mac_address,omitempty"`
		PresenceID                      string      `json:"presence_id,omitempty"`
		SuppressUnregisterNotifications bool        `json:"suppress_unregister_notifications,omitempty"`
		Timezone                        string      `json:"timezone,omitempty"`
		SIP                             SIP         `json:"sip,omitempty"`
		CallForward                     CallForward `json:"call_forward,omitempty"`
		Media                           Media       `json:"media,omitempty"`
	}

	SIP struct {
		Username                string `json:"username,omitempty"`
		StaticRoute             string `json:"static_route,omitempty"`
		Route                   string `json:"route,omitempty"`
		Realm                   string `json:"realm,omitempty"`
		Password                string `json:"password,omitempty"`
		Number                  string `json:"number,omitempty"`
		Method                  string `json:"method,omitempty"` //password or IP
		IP                      string `json:"ip,omitempty"`
		InviteFormat            string `json:"invite_forma,omitemptyt"` //npan 1npan e.164
		IgnoreCompleteElsewhere bool   `json:"ignore_complete_elsewhere,omitempty"`
		ExpireSeconds           int    `json:"expire_seconds,omitempty"`
	}

	Media struct {
		FaxOption        bool       `json:"fax_option,omitempty"`
		ProgressTimeout  int        `json:"progress_timeout,omitempty"`
		BypassMedia      bool       `json:"bypass_media,omitempty"`
		IgnoreEarlyMedia bool       `json:"ignore_early_media,omitempty"`
		Encryption       Encryption `json:"encryption,omitempty"`
		Audio            Audio      `json:"audio,omitempty"`
		Video            Video      `json:"video,omitempty"`
	}

	Encryption struct {
		EnforceSecurity bool     `json:"enforce_security,omitempty"`
		Methods         []string `json:"methods,omitempty"`
	}

	Audio struct {
		Codecs []string `json:"codecs,omitempty"`
	}

	Video struct {
		Codecs []string `json:"codecs,omitempty"`
	}
)

func (api *DevicesAPIService) CreateDevice(ctx context.Context, acc string, input *Device) (dev *Device, err error) {
	var response struct {
		Data Device `json:"data"`
		ResponseEnvelope
	}

	if acc == "" {
		return nil, reportError("account id is required field")
	}

	if input.Name == "" {
		return nil, reportError("name of the device is required field")
	}

	params := Request{
		CTX:    ctx,
		Method: "PUT",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/devices",
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

	dev = &response.Data

	return dev, nil
}

func (api *DevicesAPIService) ListDevices(ctx context.Context, acc string, disablePagination bool) (devices []Device, err error) {
	var response struct {
		Data []Device `json:"data"`
		ResponseEnvelope
	}

	params := Request{
		CTX:    ctx,
		Method: "GET",
		Path:   api.client.cfg.BasePath + "/accounts/" + acc + "/devices",
	}

	if disablePagination {
		params.QueryParams = map[string][]string{"paginate": []string{"false"}}
	}

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
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, reportError("Status: %v, Body: %s", resp.Status, bodyBytes)
	}

	decoder := json.NewDecoder(resp.Body)

	decErr := decoder.Decode(&response)
	if decErr != nil {
		return nil, reportError("can't decode response: %v", decErr)
	}

	devices = response.Data

	return devices, nil
}
