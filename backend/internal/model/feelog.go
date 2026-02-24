package model

import "time"

// Feelog is a model for a feelog
type Feelog struct {
    ID        int       `json:"id"`
    FeelScore int       `json:"feel_score"`
    Note      string    `json:"note"`
    CreatedAt time.Time `json:"created_at"`
}