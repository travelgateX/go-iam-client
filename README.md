# go-iam-client
GraphQL Go client to connect with the TravelgateX IAM API

* IAM API strong Go types for response data
* Build and execute any IAM API request
* Help package with some basic requests
* Options to log and debug transactions

## Installation
Make sure you have a working Go environment. To install go-iam-client, simply run:

```
$ go get github.com/travelgateX/go-iam-client
```

## Initialization
The service endpoint can be provide or you can let the library choose the endpoint depending on the environment variable DEPLOY_MODE. Also you must provide a valid TravelgateX bearer with permission to manage the Iam-API. There are two constructors:
```go
// IAM client, endpoint provided
iamController := iam.NewClient("Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImt...", "https://api...")

// IAM default client, endpoint chosen by DEPLOY_MODE environment variable
iamController := iam.NewDefaultClient("Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImt...")
```

## Usage
```go
// Execute Find all organizations 
res, err := i.Organizations()

// Execute find organizations by codes. Codes is a slices of strings
codes := ["org1", "org2"]
res, err :=  i.OrganizationsByCode(codes)

// Execute updateGroup, needed api to add or delete and the group to update
res, err := i.updateGroups("iam", "gr1", "ADD")

// Execute new customized query 
res, err := iamController.NewQuery(`
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
	`)


// Log errors
if err != nil {
    log.Fatal(err)
}

// Debug Mode to log transactions
// set 'true' before run a transaction
ent.DebugMode(true)
res, err := iamController.Organizations()

```

## Road Map
Interesting features to develop:

* Add unit tests
* Synchronization of the data model from the endpoint
* More request templates of the most used functions
* Possibility of obtaining raw response without fit the result with the data model

## Contribution
1. Fork ( https://github.com/[my-github-username]/go-iam-client/fork )
2. Create new feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push branch (`git push origin my-new-feature`)
5. Create new Pull Request
