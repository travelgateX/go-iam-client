package main

import (
	"log"

	"github.com/travelgateX/go-iam-client/iam"
	"github.com/travelgateX/go-iam-client/model"
)

func main() {

	// Impersonate example
	i := iam.NewDefaultClient("BEARER")
	i.DebugMode(true)

	inputCreateOrg := model.CreateOrganizationInput{
		Organization: "pikachu",
		User:         "slucea@travelgatex.com",
	}
	res, err := i.CreateOrganization(inputCreateOrg)
	if err != nil {
		log.Printf("ERROR: %v", err)
	}

	log.Printf("QUERY : %v", res.Mutation.CreateOrganization.Code)
}
