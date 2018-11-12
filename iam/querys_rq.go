package iam

import (
	"strconv"
	"strings"
)

// accessesRQ returns graphql request query
func organizationRQ() string {
	rq := `
		query{
			admin{
				organizations{
					edges{
						node{
							code
						}
					}
				}
			}
		}
	`
	return rq
}