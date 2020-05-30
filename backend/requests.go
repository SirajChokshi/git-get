package main

import (
	"context"
	"os"

	"github.com/machinebox/graphql"
)

// API URL (GitHub GraphQL v4)
const API_URL = "https://api.github.com/graphql"

func makeRequest(username string) interface{} {

	graphqlClient := graphql.NewClient(API_URL)

	type User struct {
		name          string
		avatarUrl     string
		websiteUrl    string
		followers     map[string]int
		following     map[string]int
		location      string
		createdAt     string
		company       string
		bio           string
		email         string
		organizations map[string]map[string]string
		repositories  struct {
			nodes struct {
				name            string
				primaryLanguage map[string]string
				languages       map[string]struct {
					size string
					node map[string]string
				}
			}
		}
	}

	/*
	 *	Queries:
	 *	- User profile information
	 *	- 6 Organizations
	 *	- 50 Repositories
	 * 		- Most Used language
	 *		- 6 Languages
	 *		- 20 Collaborators
	 *	-
	 */

	graphqlRequest := graphql.NewRequest(`
        query {
            user(login: "` + username + `") {
				name,
				avatarUrl,
				websiteUrl,
				followers {
					totalCount
				},
				following {
					totalCount
				},
				location,
				createdAt,
				company,
				bio,
				email,
				organizations (first: 6) {
					nodes {
						login
					}
				},
				repositories(first: 50) {
					nodes {
					  name,
					  primaryLanguage {
						name
					  },
					  languages(first: 6) {
						edges {
						  size,
						  node {
							name
						  }
						}
					  },
					}
				  }
            }
        }
	`)

	// Set Authorization Header
	graphqlRequest.Header.Set("Authorization", "Bearer "+os.Getenv("GIT_GET_TOKEN"))

	var graphqlResponse User

	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}

	return graphqlResponse
}

/*
 *	returns the number of requests remaining on the token as an integer, on error returns -1
 *	note this endpoint does not count against the rate limit
 */
func getRemainingRequests() int {
	graphqlClient := graphql.NewClient(API_URL)
	graphqlRequest := graphql.NewRequest(`
        query {
            rateLimit {
				remaining
			}
        }
	`)

	type response map[string]map[string]int

	// Set Authorization Header
	graphqlRequest.Header.Set("Authorization", "Bearer "+os.Getenv("GIT_GET_TOKEN"))
	var graphqlResponse response
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		return -1
	}

	return graphqlResponse["rateLimit"]["remaining"]
}

//TODO: Implement Collaborators Query with Github REST v3 API
// collaborators(first: 10, affiliation: ALL) {
// 	edges {
// 	  node {
// 		  login,
// 		  avatarUrl
// 		}
// 	}
// }
