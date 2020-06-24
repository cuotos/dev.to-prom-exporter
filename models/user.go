package models

type User struct {
	TypeOf          string      `json:"type_of"`
	ID              int         `json:"id"`
	Username        string      `json:"username"`
	Name            string      `json:"name"`
	Summary         string      `json:"summary"`
	TwitterUsername string      `json:"twitter_username"`
	GithubUsername  string      `json:"github_username"`
	WebsiteURL      interface{} `json:"website_url"`
	Location        string      `json:"location"`
	JoinedAt        string      `json:"joined_at"`
	ProfileImage    string      `json:"profile_image"`
}
