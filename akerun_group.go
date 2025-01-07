package akerun

import (
	"context"
	"net/http"
	"net/url"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const apiPathAkerunGroup = "akerun_groups"

type AkerunGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Memo string `json:"memo"`
}

type AkerunGroupDetailed struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Memo    string `json:"memo"`
	Akeruns []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		ImageURL string `json:"image_url"`
	} `json:"akeruns"`
}

type AkerunGroupList struct {
	AkerunGroups []AkerunGroup `json:"akerun_groups"`
}

type akerunGroupRow struct {
	AkerunGroup AkerunGroup `json:"akerun_group"`
}
type akerunGroupDetailedRow struct {
	AkerunGroup AkerunGroupDetailed `json:"akerun_group"`
}

func (c *Client) GetAkerunGroups(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	organizationId string,
) (*AkerunGroupList, error) {
	var result AkerunGroupList
	err := c.callVersion(ctx, path.Join(apiPathOrganizations, organizationId, apiPathAkerunGroup), http.MethodGet, oauth2Token, nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetAkerunGroup(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	organizationId string,
	akerunGroupId string,
) (*AkerunGroupDetailed, error) {
	var result akerunGroupDetailedRow
	err := c.callVersion(ctx, path.Join(apiPathOrganizations, organizationId, apiPathAkerunGroup, akerunGroupId), http.MethodGet, oauth2Token, nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result.AkerunGroup, nil
}

type AkerunGroupCreateParameter struct {
	Name string `url:"name"`
	Memo string `url:"memo"`
}

func (c *Client) CreateAkerunGroup(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	organizationId string,
	params AkerunGroupCreateParameter,
) (*AkerunGroup, error) {
	var result akerunGroupRow
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	err = c.callVersion(ctx, path.Join(apiPathOrganizations, organizationId, apiPathAkerunGroup), http.MethodPost, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result.AkerunGroup, nil
}

type AkerunGroupUpdateParameter struct {
	Name string `url:"name"`
	Memo string `url:"memo"`
}

func (c *Client) UpdateAkerunGroup(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	organizationId string,
	akerunGroupId string,
	params AkerunGroupUpdateParameter,
) (*AkerunGroup, error) {
	var result akerunGroupRow
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	err = c.callVersion(ctx, path.Join(apiPathOrganizations, organizationId, apiPathAkerunGroup, akerunGroupId), http.MethodPut, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result.AkerunGroup, nil
}

func (c *Client) DeleteAkerunGroup(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	organizationId string,
	akerunGroupId string,
) error {
	err := c.callVersion(ctx, path.Join(apiPathOrganizations, organizationId, apiPathAkerunGroup, akerunGroupId), http.MethodDelete, oauth2Token, nil, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AddAkerunToGroup(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	organizationId string,
	akerunGroupId string,
	akerunIds ...string,
) error {
	var result akerunGroupDetailedRow
	v := url.Values{}
	for _, id := range akerunIds {
		v.Add("akerun_ids[]", id)
	}

	err := c.callVersion(ctx, path.Join(apiPathOrganizations, organizationId, apiPathAkerunGroup, akerunGroupId, "akeruns"), http.MethodPost, oauth2Token, v, nil, &result)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) RemoveAkerunFromGroup(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	organizationId string,
	akerunGroupId string,
	akerunIds ...string,
) error {
	var result akerunGroupDetailedRow
	v := url.Values{}
	for _, id := range akerunIds {
		v.Add("akerun_ids[]", id)
	}

	err := c.callVersion(ctx, path.Join(apiPathOrganizations, organizationId, apiPathAkerunGroup, akerunGroupId, "akeruns"), http.MethodDelete, oauth2Token, v, nil, &result)
	if err != nil {
		return err
	}
	return nil
}
