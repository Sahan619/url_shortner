package models

import "time"

type URL struct {
	ID              int        `json:"id"`
	OriginalURL     string     `json:"originalUrl"`
	ShortCode       string     `json:"shortCode"`
	ShortURL        string     `json:"shortUrl"`
	Clicks          int        `json:"clicks"`
	CreatedAt       time.Time  `json:"createdAt"`
	ManagementToken string     `json:"-"`
	ExpiresAt       *time.Time `json:"expiresAt,omitempty"`
}
