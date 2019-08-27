//This module implements functions of the Channels API
//you may find documentation here: https://github.com/2600hz/kazoo/blob/master/applications/crossbar/doc/channels.m

package kazooapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

var (
	_ context.Context
)

type ChannelsAPIService service

type Channel struct {
	Answered        bool      `json:"answered"`
	AuthorizingID   string    `json:"authorizing_id"`
	AuthorizingType string    `json:"authorizing_ype"`
	Destination     string    `json:"destination"`
	Direction       string    `json:"direction"`
	OtherLeg        string    `json:"other_leg"`
	OwnerID         string    `json:"owner_id"`
	PresenceID      string    `json:"presence_id"`
	Timestamp       time.Time `json:"timestamp"`
	Username        string    `json:"username"`
	UUID            string    `json:"uuid"`
}

type ChannelList struct {
	List []Channel
}

//ListGlobalChannels returns a global list of channels
//for the whole cluster
//It should explicitely enabled by an admin
//system_config->crossbar.channels->system_wide_channels_list = true
func (chanapi *ChannelsAPIService) ListGlobalChannels(ctx context.Context) (chl []Channel, err error) {
		params := Request{
			CTX:    ctx,
			Method: "GET",
			Path:   chanapi.client.cfg.BasePath + "/channels",
		}

		req, err := chanapi.client.prepareRequest(&params)
		if err != nil {
			return nil, reportError("Can't prepare a request %s", err)
		}

		resp, err := chanapi.client.callAPI(ctx, req)
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
			return nil, reportError("Can't decode response: %v", decErr)
		}

		fmt.Println(&env.Data)

		return chl, err

}

func (chanapi *ChannelsAPIService) ListAccountChannels(ctx context.Context, acc string) (chl []Channel, err error) {
	params := Request{
		CTX:    ctx,
		Method: "GET",
		Path:   chanapi.client.cfg.BasePath + "/accounts/" + acc + "/channels",
	}

	req, err := chanapi.client.prepareRequest(&params)
	if err != nil {
		return nil, reportError("Can't prepare a request %s", err)
	}

	resp, err := chanapi.client.callAPI(ctx, req)
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
		return nil, reportError("Can't decode response: %v", decErr)
	}

	fmt.Println(&env.Data)

	return chl, err

}
