package iam

import "strings"

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

func organizationsByCodeRQ(codes []string) string {
	rq := `
	{
		admin{
		  organizations(codes: "$CODES$"){
			edges{
			  node{
				code
			  }
			}
		  }
		}
	  }
	`
	codeFilter := sliceToQuotedStringFormat(codes)
	rq = strings.Replace(rq, "$CODES$", codeFilter, 1)
	return rq
}
