//This module implements functions of the Channels API
//you may find documentation here: https://github.com/2600hz/kazoo/blob/master/applications/crossbar/doc/channels.m

package kazooapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
)

var (
	_ context.Context
)

//RecordingsAPIService represents API for tangling recordings
type RecordingsAPIService service

type Recording struct {
	ID                string                 `json:"id"`
	Description       string                 `json:"description"`
	Direction         string                 `json:"direction"`
	Duration          int                    `json:"duration"`
	DurationMS        int                    `json:"duration_ms"`
	CallID            string                 `json:"call_id"`
	CalleeIDName      string                 `json:"callee_id_name"`
	CalleeIDNumber    string                 `json:"callee_id_number"`
	CallerIDName      string                 `json:"caller_id_name"`
	CallerIDNumber    string                 `json:"caller_id_number"`
	CdrID             string                 `json:"cdr_id"`
	MediaSource       string                 `json:"media_source"`
	MediaType         string                 `json:"media_type"`
	ContentType       string                 `json:"content_type"`
	Name              string                 `json:"name"`
	From              string                 `json:"from"`
	To                string                 `json:"to"`
	InteractionID     string                 `json:"interaction_id"`
	OwnerID           string                 `json:"owner_id"`
	Request           string                 `json:"request"`
	Start             int                    `json:"start"`
	SourceType        string                 `json:"source_type"`
	CustomChannelVars map[string]interface{} `json:"custom_channel_vars"`
}

type RecordingsList struct {
	List []Recording
}

//ListRecordings returns a list of recordings for the account
func (recapi *RecordingsAPIService) ListRecordings(ctx context.Context, acc string) (rec []Recording, err error) {
	params := Request{
		CTX:    ctx,
		Method: "GET",
		Path:   recapi.client.cfg.BasePath + "/accounts/" + acc + "/recordings",
	}

	req, err := recapi.client.prepareRequest(&params)
	if err != nil {
		return nil, reportError("Can't prepare a request %s", err)
	}

	resp, err := recapi.client.callAPI(ctx, req)
	if err != nil || resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, reportError("Status: %v, Body: %s", resp.Status, bodyBytes)
	}

	env := RequestEnvelope{}

	decoder := json.NewDecoder(resp.Body)

	decErr := decoder.Decode(&env)
	if decErr != nil {
		return nil, reportError("Can't decode response: %s", decErr)
	}

	return rec, err

}

func (recapi *RecordingsAPIService) GetRecording(ctx context.Context, acc, recording string) (rec *Recording, err error) {
	rec = &Recording{}

	params := Request{
		CTX:    ctx,
		Method: "GET",
		Path:   recapi.client.cfg.BasePath + "/accounts/" + acc + "/recordings/" + recording,
	}

	req, err := recapi.client.prepareRequest(&params)
	if err != nil {
		return nil, reportError("can't prepare a request %s", err)
	}

	resp, err := recapi.client.callAPI(ctx, req)
	if err != nil || resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, reportError("StatusCode: %v, Status: %v, Body: %s", resp.StatusCode, resp.Status, bodyBytes)
	}

	env := RequestEnvelope{}
	env.Data = rec

	decoder := json.NewDecoder(resp.Body)

	decErr := decoder.Decode(&env)
	if decErr != nil {
		return nil, reportError("Can't decode response: %v", decErr)
	}

	return rec, err

}
