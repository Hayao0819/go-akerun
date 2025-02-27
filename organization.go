package akerun

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	apiPathOrganizations = "organizations"
)

// id represents an ID of an organization.
type id struct {
	ID string `json:"id"`
}

// OrganizationList represents a list of organizations in Akerun API.
type OrganizationList struct {
	Organizations []id `json:"organizations"`
}

// organizationRow represents a row in the organization table.
type organizationRow struct {
	Organization Organization `json:"organization"`
}

// Organization represents the detailed information of an organization.
type Organization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OrganizationsParameter represents the parameters for GetOrganizations method.
type OrganizationsParameter struct {
	Limit    uint32 `url:"limit,omitempty"`
	IdAfter  string `url:"id_after,omitempty"`
	IdBefore string `url:"id_before,omitempty"`
}

// GetOrganizations returns a list of organizations.
func (c *Client) GetOrganizations(
	ctx context.Context, oauth2Token *oauth2.Token, params OrganizationsParameter) (*OrganizationList, error) {

	var result OrganizationList
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	err = c.callVersion(ctx, apiPathOrganizations, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetOrganization retrieves the details of an organization with the specified ID.
func (c *Client) GetOrganization(ctx context.Context, oauth2Token *oauth2.Token, id string) (*Organization, error) {
	var result organizationRow
	err := c.callVersion(ctx, path.Join(apiPathOrganizations, id), http.MethodGet, oauth2Token, nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result.Organization, nil
}
