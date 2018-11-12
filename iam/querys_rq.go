package iam

// accessesRQ returns graphql request query
func OrganizationRQ() string {
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
