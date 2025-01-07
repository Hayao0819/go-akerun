package akerun

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	apiPathKeys = "keys"
)

type Key struct {
	ID                string `json:"id"`
	Role              string `json:"role"`
	ScheduleType      string `json:"schedule_type"`
	TemporarySchedule struct {
		StartDateTime string `json:"start_datetime"`
		EndDateTime   string `json:"end_datetime"`
	} `json:"temporary_schedule"`
	RecurringSchedule struct {
		DaysOfWeek []uint32 `json:"days_of_week"`
		StartTime  string   `json:"start_time"`
		EndTime    string   `json:"end_time"`
	} `json:"recurring_schedule"`
	Keys struct {
		KeyUrl            string `json:"key_url"`
		PasswordProtected bool   `json:"password_protected"`
	} `json:"keys"`
	Akerun struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"akerun"`
	User struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
}

type KeysList struct {
	Keys []Key `json:"keys"`
}

type keyRow struct {
	Key Key `json:"key"`
}

type KeysParameter struct {
	UserId   string `url:"user_id,omitempty"`
	AkerunId string `url:"akerun_id,omitempty"`
	Limit    uint32 `url:"limit,omitempty"`
	IdAfter  string `url:"id_after,omitempty"`
	IdBefore string `url:"id_before,omitempty"`
}

func (c *Client) GetKeys(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	organizationId string,
	params KeysParameter) (*KeysList, error) {

	var result KeysList
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	err = c.callVersion(ctx, path.Join(apiPathOrganizations, organizationId, apiPathKeys), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type KeyParameter struct {
	KeyId string `url:"key_id,omitempty"`
}

func (c *Client) GetKey(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	organizationId string,
	params KeyParameter) (*Key, error) {

	var result keyRow
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	err = c.callVersion(ctx, path.Join(apiPathOrganizations, organizationId, apiPathKeys), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result.Key, nil
}

type CreateKeyParameter struct {
	ScheduleType      string `url:"schedule_type,omitempty"`
	TemporarySchedule struct {
		StartDateTime string `url:"start_datetime,omitempty"`
		EndDateTime   string `url:"end_datetime,omitempty"`
	} `url:"temporary_schedule,omitempty"`
	RecurringSchedule struct {
		DaysOfWeek []uint32 `url:"days_of_week,omitempty"`
		StartTime  string   `url:"start_time,omitempty"`
		EndTime    string   `url:"end_time,omitempty"`
	} `url:"recurring_schedule,omitempty"`
	EnableKeyUrl   bool   `url:"enable_key_url,omitempty"`
	KeyUrlPassword string `url:"key_url_password,omitempty"`
	Role           string `url:"role,omitempty"`
}

func (c *Client) CreateKey(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	organizationId string,
	userId string,
	akerunId string,
	params CreateKeyParameter,
) (*Key, error) {
	var result keyRow
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	v.Add("user_id", userId)
	v.Add("akerun_id", akerunId)
	err = c.callVersion(ctx, path.Join(apiPathOrganizations, organizationId, apiPathKeys), http.MethodPost, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result.Key, nil
}
