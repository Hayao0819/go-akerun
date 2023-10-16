package akerun

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathOrganizations = "organizations"
)

type Organization struct {
	ID string `json:"id"`
}

type Organizations struct {
	Organizations []Organization `json:"organizations"`
}

type OrganizationsParams struct {
	Limit    uint32 `url:"limit,omitempty"`
	IdAfter  string `url:"id_after,omitempty"`
	IdBefore string `url:"id_before,omitempty"`
}

func (c *Client) GetOrganizations(
	ctx context.Context, oauth2Token *oauth2.Token, params OrganizationsParams) (*Organizations, error) {

	var result Organizations
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	err = c.call(ctx, APIPathOrganizations, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}