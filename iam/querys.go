package iam

import (
	"errors"

	"github.com/travelgateX/go-entities-client/model"
)

// Organizations IAM API query function
func (c *Client) Organizations(id int) (model.AdminQuery, error) {
	return c.NewQuery(accessesRQ(id))
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

// IAM IAM API query function
func (c *Client) IAM() (model.AdminQuery, error) {
	return model.AdminQuery{}, errors.New("not implemented")
}
