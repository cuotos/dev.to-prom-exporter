package models

import "time"

type Article struct {
	TypeOf               string    `json:"type_of"`
	ID                   int       `json:"id"`
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	CoverImage           string    `json:"cover_image"`
	Published            bool      `json:"published"`
	PublishedAt          time.Time `json:"published_at"`
	TagList              []string  `json:"tag_list"`
	Slug                 string    `json:"slug"`
	Path                 string    `json:"path"`
	URL                  string    `json:"url"`
	CanonicalURL         string    `json:"canonical_url"`
	CommentsCount        int       `json:"comments_count"`
	PublicReactionsCount int       `json:"public_reactions_count"`
	PageViewsCount       int       `json:"page_views_count"`
	PublishedTimestamp   time.Time `json:"published_timestamp"`
	BodyMarkdown         string    `json:"body_markdown"`
	User                 struct {
		Name            string `json:"name"`
		Username        string `json:"username"`
		TwitterUsername string `json:"twitter_username"`
		GithubUsername  string `json:"github_username"`
		WebsiteURL      string `json:"website_url"`
		ProfileImage    string `json:"profile_image"`
		ProfileImage90  string `json:"profile_image_90"`
	} `json:"user"`
	Organization struct {
		Name           string `json:"name"`
		Username       string `json:"username"`
		Slug           string `json:"slug"`
		ProfileImage   string `json:"profile_image"`
		ProfileImage90 string `json:"profile_image_90"`
	} `json:"organization"`
	FlareTag struct {
		Name         string `json:"name"`
		BgColorHex   string `json:"bg_color_hex"`
		TextColorHex string `json:"text_color_hex"`
	} `json:"flare_tag"`
}