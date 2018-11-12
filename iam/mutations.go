package iam

import (
	"errors"

	"github.com/travelgateX/go-entities-client/model"
)

// CreateAccess Entities API query function
func (c *Client) CreateAccess() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// UpdateAccess Entities API query function
func (c *Client) UpdateAccess() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// GrantAccessToGroup Entities API mutation function
func (c *Client) GrantAccessToGroup(id int, groups []string) (model.AdminMutation, error) {
	return c.NewMutation(grantAccessToGroupRQ(id, groups))
}

// DeleteAccessFromGroup Entities API mutation function
func (c *Client) DeleteAccessFromGroup() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// GrantSupplierToGroup Entities API mutation function
func (c *Client) GrantSupplierToGroup() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// DeleteSupplierFromGroup Entities API mutation function
func (c *Client) DeleteSupplierFromGroup() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// GrantClientToGroup Entities API mutation function
func (c *Client) GrantClientToGroup() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// DeleteClientFromGroup Entities API mutation function
func (c *Client) DeleteClientFromGroup() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// CreateClient Entities API mutation function
func (c *Client) CreateClient() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// UpdateClient Entities API mutation function
func (c *Client) UpdateClient() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// CreateProfile Entities API mutation function
func (c *Client) CreateProfile() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// UpdateProfile Entities API mutation function
func (c *Client) UpdateProfile() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// AddEntitiesToProfile Entities API mutation function
func (c *Client) AddEntitiesToProfile() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// RemoveEntitiesFromProfile Entities API mutation function
func (c *Client) RemoveEntitiesFromProfile() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}
