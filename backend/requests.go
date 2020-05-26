package main

import (
	"context"
	"fmt"
	"os"

	"github.com/machinebox/graphql"
)

// API URL (GitHub GraphQL v4)
const API_URL = "https://api.github.com/graphql"

func makeRequest(username string) interface{} {

	graphqlClient := graphql.NewClient(API_URL)

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

	var graphqlResponse interface{}

	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}

	return graphqlResponse
}

func getRateLimit() {
	graphqlClient := graphql.NewClient(API_URL)
	graphqlRequest := graphql.NewRequest(`
        query {
            rateLimit {
				remaining
			}
        }
	`)

	// Set Authorization Header
	graphqlRequest.Header.Set("Authorization", "Bearer "+os.Getenv("GIT_GET_TOKEN"))
	var graphqlResponse interface{}
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}

	fmt.Println(graphqlResponse)
}

// TODO: Implement Collaborators Query (Throws Error in current query)
// collaborators(first: 10, affiliation: ALL) {
// 	edges {
// 	  node {
// 		  login,
// 		  avatarUrl
// 		}
// 	}
// }
