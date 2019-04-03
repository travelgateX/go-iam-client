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
		  organizations(codes: $CODES$){
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

func impersonateJWT(member string) string {
	rq := `
	query{
		admin{
			members(codes:"$CODE$"){
				edges{
					node{
						memberData{
							code
							impersonationJWT{
								token
							}
						}
					}
				}
			}
		}
    
	}
	`
	rq = strings.Replace(rq, "$CODE$", member, 1)
	return rq
}

func groupsByCodeRQ(codes []string) string {
	rq := `
		query{
			admin{
				groups(codes:$CODES$){
					edges{
						node{
							groupData{
								id
								code
								type
								label
							}
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
