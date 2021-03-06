// Package iam is a tool to connect to the "IAM API Graphql" endpoint.
// There are functions that it provides basic requests for common use or you
// can use the NewQuery/NewMutation main functions.
//
// WARNING: 'IAMAPI' is a Graphql server, requests can not be typed so if you
// need more personalized request use the client Query/Mutation main function.
//
// Example of templates use:
//
// 		iamContreller := iam.NewDefaultClient("Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1N...")
// 		res, _ := iamContreller.Organizations()
// 		fmt.Printf("Access.name = %v", res.Query.Accesses.Edges[0].Node.AccessData.Name)
//
// Example of basic use:
//
// 		iamContreller := iam.NewDefaultClient("Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1N...")
// 		res, _ := iamContreller.NewQuery(`query{admin{accesses(...) ...`)
// 		fmt.Printf("Access.name = %v", res.Query.Accesses.Edges[0].Node.AccessData.Name)
//
package iam

import (
	"context"
	"log"
	"os"

	"github.com/machinebox/graphql"

	"github.com/travelgateX/go-iam-client/model"
)

// IAM API end points
const (
	IAMEndPointProd = "https://api-iam.travelgatex.com/controller/query"
	IAMEndPointDev  = "https://dev-api-iam.travelgatex.com/controller/xquery"
)

// Client : Grapqhql client
type Client struct {
	graphql *graphql.Client // graphql client
	bearer  string          // authentification bearer
}

// NewClient constructor
func NewClient(bearer, endpoint string) Client {
	cli := graphql.NewClient(endpoint)
	return Client{graphql: cli, bearer: bearer}
}

// NewDefaultClient default constructor
func NewDefaultClient(bearer string) Client {
	cli := NewClient(bearer, IAMEndPointDev)
	if os.Getenv("DEPLOY_MODE") == "prod" || os.Getenv("DEPLOY_MODE") == "localProd" {
		cli = NewClient(bearer, IAMEndPointProd)
	}
	return cli
}

// DebugMode set debug mode
func (c *Client) DebugMode(debug bool) {
	if debug {
		c.graphql.Log = func(s string) { log.Println(s) }
	} else {
		c.graphql.Log = func(s string) {}
	}
}

// NewQuery excute new graphql query
func (c *Client) NewQuery(rq string) (model.AdminQuery, error) {
	ctx := context.Background()
	req := graphql.NewRequest(rq)
	req.Header.Add("Authorization", c.bearer)

	var res model.AdminQuery
	if err := c.graphql.Run(ctx, req, &res); err != nil {
		return model.AdminQuery{}, err
	}
	return res, nil
}

// NewMutation execute new graphql mutation
func (c *Client) NewMutation(rq string) (model.AdminMutation, error) {
	ctx := context.Background()
	req := graphql.NewRequest(rq)
	req.Header.Add("Authorization", c.bearer)

	var res model.AdminMutation
	if err := c.graphql.Run(ctx, req, &res); err != nil {
		return model.AdminMutation{}, err
	}
	return res, nil
}
