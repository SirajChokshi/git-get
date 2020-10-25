export interface User {
    AvatarURL: string;
    Bio: string;
    CommitsPerDay: Commit[];
    CreatedAt: string;
    Company: string;
    Email: string;
    Followers: number;
    Following: number;
    Location: string;
    Login: string;
    Name: string;
    Organizations: string[];
    PullRequestsMade: number;
    PullRequestsReviewed: number;
    Repositories: Repository[];
    Stars: number;
    TotalContributions: number;
    WebsiteURL: string;
}

export interface Commit {
    Count: number;
    Date: string;
    Weekday: number;
}

export interface Repository {
    Commits: number;
    Name: string;
    PrimaryLanguage: string;
    Stars: number;
    Collaborators: string[];
    Languages: object;
}