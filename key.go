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
		DaysOfWeek []int  `json:"days_of_week"`
		StartTime  string `json:"start_time"`
		EndTime    string `json:"end_time"`
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

type KeysParameter struct {
	KeyId string `url:"key_id,omitempty"`
}

type keyRow struct {
	Key Key `json:"key"`
}

func (c *Client) GetKeys(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	organizationId string,
	params KeysParameter) (*Key, error) {

	var result keyRow
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	err = c.callVersion(ctx, path.Join(apiPathOrganizations), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result.Key, nil
}
