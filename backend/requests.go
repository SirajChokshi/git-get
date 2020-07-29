package main

import (
	"context"
	"encoding/json"
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
		Login         string `json:"login"`
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
				Login     string `json:"login"`
				AvatarURL string `json:"avatarUrl"`
			}
		}
		Repositories struct {
			Nodes []struct {
				Object struct {
					History struct {
						TotalCount int `json:"totalCount"`
						Nodes      []struct {
							Author struct {
								User struct {
									Login string `json:"login"`
								} `json:"user"`
							} `json:"author"`
						} `json:"nodes"`
					} `json:"history"`
				} `json:"object"`
				Name       string `json:"name"`
				Stargazers struct {
					TotalCount int `json:"totalCount"`
				} `json:"stargazers"`
				PrimaryLanguage struct {
					Name string `json:"name"`
				} `json:"primaryLanguage"`
				Languages struct {
					Edges []struct {
						Size int `json:"size"`
						Node struct {
							Name string `json:"name"`
						} `json:"node"`
					} `json:"edges"`
				} `json:"languages"`
			} `json:"nodes"`
		} `json:"repositories"`
		ContributionsCollection struct {
			ContributionCalendar struct {
				TotalContributions int `json:"totalContributions"`
				Weeks              []struct {
					ContributionDays []struct {
						ContributionCount int    `json:"contributionCount"`
						Date              string `json:"date"`
						Weekday           int    `json:"weekday"`
					} `json:"contributionDays"`
				} `json:"weeks"`
			} `json:"contributionCalendar"`
			PullRequestReviewContributions struct {
				TotalCount int `json:"totalCount"`
			} `json:"pullRequestReviewContributions"`
			PullRequestContributions struct {
				TotalCount int `json:"totalCount"`
			} `json:"pullRequestContributions"`
		} `json:"contributionsCollection"`
	}
}

// Repository ...
type Repository struct {
	Name            string
	PrimaryLanguage string
	Languages       map[string]int
	Commits         int
	Collaborators   []string
	Stars           int
}

// Day ...
type Day struct {
	Count   int
	Date    string
	Weekday int
}

// JSONUser ... type to easily access user data after a request
type JSONUser struct {
	Name                 string
	Login                string
	AvatarURL            string
	WebsiteURL           string
	Followers            int
	Following            int
	Location             string
	CreatedAt            string
	Company              string
	Bio                  string
	Email                string
	Organizations        []string
	Repositories         []Repository
	TotalContributions   int
	Stars                int
	CommitsPerDay        []Day
	PullRequestsMade     int
	PullRequestsReviewed int
}

func makeRequest(username string) *JSONUser {

	graphqlClient := graphql.NewClient(API_URL)

	/*
	 *	Queries:
	 *	- User profile information
	 *	- 6 Organizations
	 *		- Logins
	 *		- avatarURL
	 *	- 50 Repositories
	 *		- Number of Commits
	 * 		- Most Used language
	 *		- 6 Languages
	 *		- Last 50 commit authors
	 *		- Number of Stars
	 *	- User Contributions
	 *		- Number of Contributions
	 *		- Commits per Day (of Week)
	 *		- Number of PRs Reviewed
	 *		- Number of PRs Made
	 */

	graphqlRequest := graphql.NewRequest(`
        query {
            user(login: "` + username + `") {
				name
				login
				avatarUrl
				websiteUrl
				followers {
				  totalCount
				}
				following {
				  totalCount
				}
				location
				createdAt
				company
				bio
				email
				organizations(first: 6) {
				  nodes {
					login
					avatarUrl
				  }
				}
				repositories(first: 100) {
				  nodes {
					object(expression: "master") {
					  ... on Commit {
						history(first: 100) {
						  totalCount
						  nodes {
							author {
							  user {
								login
							  }
							}
						  }
						}
					  }
					}
					name
					stargazers {
					  totalCount
					}
					primaryLanguage {
					  name
					}
					languages(first: 6) {
					  edges {
						size
						node {
						  name
						}
					  }
					}
				  }
				}
				contributionsCollection {
					contributionCalendar {
					  totalContributions
					  weeks {
						contributionDays {
						  contributionCount
						  date
						  weekday
						}
					  }
					}
					pullRequestReviewContributions {
					  totalCount
					}
					pullRequestContributions {
					  totalCount
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

	// fmt.Println(graphqlResponse)

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
		out.PrimaryLanguage = repo.PrimaryLanguage.Name

		languages := make(map[string]int)

		for _, lang := range repo.Languages.Edges {
			languages[lang.Node.Name] = lang.Size
		}

		out.Languages = languages

		out.Commits = repo.Object.History.TotalCount

		contributors := make(map[string]bool)

		for _, node := range repo.Object.History.Nodes {
			if node.Author.User.Login != res.Login {
				contributors[node.Author.User.Login] = true
			}
		}

		var collaborators []string

		for username := range contributors {
			collaborators = append(collaborators, username)
		}

		out.Collaborators = collaborators
		out.Stars = repo.Stargazers.TotalCount

		repos = append(repos, out)
	}

	totalStars := 0

	for _, repo := range repos {
		totalStars += repo.Stars
	}

	var days []Day

	for _, weeks := range res.ContributionsCollection.ContributionCalendar.Weeks {
		for _, day := range weeks.ContributionDays {
			var current Day
			current.Date = day.Date
			current.Weekday = day.Weekday
			current.Count = day.ContributionCount
			days = append(days, current)
		}
	}

	formattedUser := JSONUser{
		Name:                 res.Name,
		Login:                res.Login,
		AvatarURL:            res.AvatarURL,
		WebsiteURL:           res.WebsiteURL,
		Followers:            res.Followers["totalCount"],
		Following:            res.Following["totalCount"],
		Location:             res.Location,
		CreatedAt:            res.CreatedAt,
		Company:              res.Company,
		Bio:                  res.Bio,
		Email:                res.Email,
		Organizations:        organizations,
		Repositories:         repos,
		CommitsPerDay:        days,
		Stars:                totalStars,
		TotalContributions:   res.ContributionsCollection.ContributionCalendar.TotalContributions,
		PullRequestsMade:     res.ContributionsCollection.PullRequestContributions.TotalCount,
		PullRequestsReviewed: res.ContributionsCollection.PullRequestReviewContributions.TotalCount,
	}

	// fmt.Println(formattedUser)

	return &formattedUser
}

/*
 *	returns the number of requests remaining on the token as an integer, on error returns -1
 *	note this endpoint does not count against the rate limit
 */
func getRemainingRequests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
