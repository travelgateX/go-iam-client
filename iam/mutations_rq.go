package iam

import (
	"strings"
)

func updateGroupsRQ(api string, group string, method string) string {

	rq := `
	mutation{
		admin{
		  updateGroup(
			group:{
			  api:"$API$" 
			  group:"$GROUP$"
			}
			method:"$METHOD$"
		  ){
			error{
			  code
			  type
			  description
			}
		  }
		}
	  }
	}
  `
	rq = strings.Replace(rq, "$API", api, 1)
	rq = strings.Replace(rq, "$GROUP$", group, 1)
	rq = strings.Replace(rq, "$METHOD$", method, 1)
	return rq
}
