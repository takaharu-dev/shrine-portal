package model

type Shrine struct {
	ID           int64
	Name         string
	Slug         string
	Description  string
	Address      string
	PrefectureID int64
	ImageURL     string
	WebsiteURL   string
	Access       string
	CreatedAt    string
	UpdatedAt    string
}
