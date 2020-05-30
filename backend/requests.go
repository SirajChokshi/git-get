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

	type Response struct {
		User struct {
			Name          string `json:"name"`
			AvatarURL     string `json:"avatarUrl"`
			WebsiteURL    string `json:"websiteUrl"`
			Followers     map[string]int
			Following     map[string]int
			Location      string `json:"location"`
			CreatedAt     string `json:"createdAt"`
			Company       string `json:"company"`
			Bio           string `json:"bio"`
			Email         string `json:"email"`
			Organizations struct {
				Nodes []struct {
					Login string `json:"login"`
				}
			}
			Repositories struct {
				Nodes []struct {
					Name            string
					PrimaryLanguage map[string]string
					Languages       map[string][]struct {
						Size int
						Node map[string]string
					}
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

	var graphqlResponse Response

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
