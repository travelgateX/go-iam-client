package main

import (
	"log"

	"go-iam-client/iam"
)

func main() {

	// Impersonate example
	i := iam.NewDefaultClient("BEARER_PLACEHOLDER")

	res, err := i.Impersonate("MEMBER_PLACEHOLDER")
	if err != nil {
		log.Printf("ERROR: %v", err)
	}

	log.Printf("%v", res.Query.Members.Edges[0].Node.MemberData.ImpersonationJWT.Token)
}
