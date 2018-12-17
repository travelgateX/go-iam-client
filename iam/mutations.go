package iam

import (
	"errors"

	"go-iam-client/model"
)

// CreateMemeber IAM API mutation function
func (c *Client) CreateMemeber(input model.CreateMemberInput) (model.AdminMutation, error) {
	return c.NewMutation(createMemberRQ(input))
}

// CreateOrganization IAM API mutation function
func (c *Client) CreateOrganization(input model.CreateOrganizationInput) (model.AdminMutation, error) {
	return c.NewMutation(createOrganizationRQ(input))
}

// DeleteOrganization IAM API mutation function
func (c *Client) DeleteOrganization(input model.DeleteOrganizationInput) (model.AdminMutation, error) {
	return c.NewMutation(deleteOrganizationRQ(input))
}

// CreateGroup IAM API mutation function
func (c *Client) CreateGroup() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// UpdateMember IAM API mutation function
func (c *Client) UpdateMember(input model.UpdateMemberInput) (model.AdminMutation, error) {
	return c.NewMutation(updateMemberRQ(input))
}

// UpdateGroups Entities API mutation function
func (c *Client) UpdateGroups(input model.UpdateGroupInput, method string) (model.AdminMutation, error) {
	return c.NewMutation(updateGroupsRQ(input, method))
}

// DeleteMember IAM API mutation function
func (c *Client) DeleteMember() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// DeleteGroup IAM API mutation function
func (c *Client) DeleteGroup() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}
