package iam

import (
	"errors"

	"github.com/travelgateX/go-iam-client/model"
)

// Organizations IAM API query function
func (c *Client) Organizations() (model.AdminQuery, error) {
	return c.NewQuery(organizationRQ())
}

// OrganizationsByCode IAM query function
func (c *Client) OrganizationsByCode(codes []string) (model.AdminQuery, error) {
	return c.NewQuery(organizationsByCodeRQ(codes))
}

// Impersonate IAM memberData function
func (c *Client) Impersonate(member string) (model.AdminQuery, error) {
	return c.NewQuery(impersonateJWT(member))
}

// GroupsByCode IAM query function
func (c *Client) GroupsByCode(codes []string) (model.AdminQuery, error) {
	return c.NewQuery(groupsByCodeRQ(codes))
}

// Products IAM API query function
func (c *Client) Products() (model.AdminQuery, error) {
	return model.AdminQuery{}, errors.New("not implemented")
}

// Members IAM API query function
func (c *Client) Members() (model.AdminQuery, error) {
	return model.AdminQuery{}, errors.New("not implemented")
}

// Groups IAM API query function
func (c *Client) Groups() (model.AdminQuery, error) {
	return model.AdminQuery{}, errors.New("not implemented")
}

// Apis IAM API query function
func (c *Client) Apis() (model.AdminQuery, error) {
	return model.AdminQuery{}, errors.New("not implemented")
}

// Resources IAM API query function
func (c *Client) Resources() (model.AdminQuery, error) {
	return model.AdminQuery{}, errors.New("not implemented")
}

// Roles IAM API query function
func (c *Client) Roles() (model.AdminQuery, error) {
	return model.AdminQuery{}, errors.New("not implemented")
}

// Operations IAM API query function
func (c *Client) Operations() (model.AdminQuery, error) {
	return model.AdminQuery{}, errors.New("not implemented")
}
