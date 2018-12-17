package iam

import (
	"log"
	"strings"

	"github.com/travelgateX/go-iam-client/model"
)

func updateGroupsRQ(input model.UpdateGroupInput, method string) string {

	rq := `
	mutation{
		admin{
		  updateGroup(
			group:{
			  api:"$API$"
			  group:"$GROUP$"
			}
			method:$METHOD$
		  ){
			code
			error{
			  code
			  type
			  description
			}
		  }
		}
	}
  `
	rq = strings.Replace(rq, "$API$", input.API, 1)
	rq = strings.Replace(rq, "$GROUP$", input.Group, 1)
	rq = strings.Replace(rq, "$METHOD$", method, 1)
	return rq
}

func createOrganizationRQ(input model.CreateOrganizationInput) string {
	rq := `
	mutation{
		admin{
			createOrganization(organization:{
				user: "$USER$"
				organization:"$ORG$" 
				template:ORGANIZATION_DEFAULT
			}){
				code
				adviseMessage{
					 code
					description
					level
				}
			}
		}
	}
	`
	rq = strings.Replace(rq, "$USER$", input.User, 1)
	rq = strings.Replace(rq, "$ORG$", input.Organization, 1)
	log.Printf("QUERY : %v", rq)

	return rq
}

func deleteOrganizationRQ(input model.DeleteOrganizationInput) string {
	rq := `
	mutation{
		admin{
			deleteOrganization(organization:{
				organization:"$ORG$"
			}){
				code
				adviseMessage{
					code
					description
					level
				}
			}
		}
	}
	`
	rq = strings.Replace(rq, "$ORG$", input.Organization, 1)
	return rq
}

func updateMemberRQ(input model.UpdateMemberInput) string {
	rq := `
	mutation{
		admin{
			updateMember(member:{
				member: $MEMBER$
				info: "$INFO$"
				group: "$GROUP$"
				resources: $RESOURCES$
				role: "$ROLE$"
				method: $METHOD$
			}) {
				code
				adviseMessage{
					code
					description
					level
				}
			}
		}
	}
	`
	rq = strings.Replace(rq, "$MEMBER$", input.Member, 1)
	rq = strings.Replace(rq, "$INFO$", input.Info, 1)
	rq = strings.Replace(rq, "$GROUP$", input.Group, 1)

	resources := sliceToQuotedStringFormat(input.Resources)
	rq = strings.Replace(rq, "$RESOURCES$", resources, 1)

	rq = strings.Replace(rq, "$ROLE$", input.Member, 1)
	rq = strings.Replace(rq, "$METHOD$", input.Method.String(), 1)

	return rq
}

func createMemberRQ(input model.CreateMemberInput) string {
	rq := `
	mutation{
		admin{
			createMember(member:{
				member:"$MEMBER$"
				info: "$INFO$"
				type:$TYPE$
				group: "$GROUP$"
				resources:$RESOURCES$
				role:"$ROLE$"
			}){
				code
				adviseMessage{
					code
					description
					level
				}
			}
		}
	}
	`
	rq = strings.Replace(rq, "$MEMBER$", input.Member, 1)
	rq = strings.Replace(rq, "$INFO$", input.Info, 1)
	rq = strings.Replace(rq, "$GROUP$", input.Group, 1)

	resources := sliceToQuotedStringFormat(input.Resources)
	rq = strings.Replace(rq, "$RESOURCES$", resources, 1)

	rq = strings.Replace(rq, "$ROLE$", input.Role, 1)
	rq = strings.Replace(rq, "$TYPE$", input.Type.String(), 1)

	log.Printf("QUERY : %v", rq)
	return rq
}
