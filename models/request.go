package models

type RequestNewStore struct {
	Name        string `json:"name" required:"json"`
	Username    string `json:"username" required:"json"`
	Description string `json:"description"`
	Type        string `json:"type" required:"json"`
	Source      string `json:"source" required:"json"`
	Image       string `json:"image"` // base 64 encode image
}
