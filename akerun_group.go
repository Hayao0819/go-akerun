package akerun

import (
	"context"
	"net/http"
	"path"

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
	var result akerunGroupRow
	err := c.callVersion(ctx, path.Join(apiPathOrganizations, organizationId, apiPathAkerunGroup, akerunGroupId), http.MethodGet, oauth2Token, nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result.AkerunGroup, nil
}
