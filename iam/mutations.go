package iam

import (
	"errors"

	"go-iam-client/model"
)

// CreateMemeber IAM API mutation function
func (c *Client) CreateMemeber() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// CreateGroup IAM API mutation function
func (c *Client) CreateGroup() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// UpdateMember IAM API mutation function
func (c *Client) UpdateMember() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// UpdateGroups Entities API mutation function
func (c *Client) UpdateGroups(api string, group string, method string) (model.AdminMutation, error) {
	return c.NewMutation(updateGroupsRQ(api, group, method))
}

// DeleteMember IAM API mutation function
func (c *Client) DeleteMember() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// DeleteGroup IAM API mutation function
func (c *Client) DeleteGroup() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}
