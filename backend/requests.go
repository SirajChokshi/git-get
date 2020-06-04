package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/machinebox/graphql"
)

// API URL (GitHub GraphQL v4)
const API_URL = "https://api.github.com/graphql"

// Response ... type to store and access JSON GraphQL response
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
				Languages       struct {
					Edges []struct {
						Size int
						Node struct {
							Name string
						}
					}
				}
			}
		}
	}
}

// Repository ...
type Repository struct {
	Name            string
	PrimaryLanguage string
	Languages       map[string]int
}

// JSONUser ... type to easily access user data after a request
type JSONUser struct {
	Name          string
	AvatarURL     string
	WebsiteURL    string
	Followers     int
	Following     int
	Location      string
	CreatedAt     string
	Company       string
	Bio           string
	Email         string
	Organizations []string
	Repositories  []Repository
}

func makeRequest(username string) *JSONUser {

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

	var graphqlResponse Response

	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}

	return formatUser(&graphqlResponse)
}

func formatUser(Data *Response) *JSONUser {
	raw := *Data
	res := raw.User

	var organizations []string
	var repos []Repository

	for _, org := range res.Organizations.Nodes {
		organizations = append(organizations, org.Login)
	}

	for _, repo := range res.Repositories.Nodes {
		var out Repository

		out.Name = repo.Name
		out.PrimaryLanguage = repo.PrimaryLanguage["name"]

		languages := make(map[string]int)

		for _, lang := range repo.Languages.Edges {
			languages[lang.Node.Name] = lang.Size
		}

		out.Languages = languages

		repos = append(repos, out)
	}

	formattedUser := JSONUser{
		Name:          res.Name,
		AvatarURL:     res.AvatarURL,
		WebsiteURL:    res.WebsiteURL,
		Followers:     res.Followers["totalCount"],
		Following:     res.Following["totalCount"],
		Location:      res.Location,
		CreatedAt:     res.CreatedAt,
		Company:       res.Company,
		Bio:           res.Bio,
		Email:         res.Email,
		Organizations: organizations,
		Repositories:  repos,
	}

	fmt.Println(formattedUser)

	return &formattedUser
}

/*
 *	returns the number of requests remaining on the token as an integer, on error returns -1
 *	note this endpoint does not count against the rate limit
 */
func getRemainingRequests(w http.ResponseWriter, r *http.Request) {
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
		json.NewEncoder(w).Encode(map[string]interface{}{"requests": 0})
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"requests": graphqlResponse["rateLimit"]["remaining"]})
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
