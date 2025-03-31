package model

import (
	"gorm.io/gorm"
)

// This is a base model that every model
// in the applicaiton should embed.
type BaseModel struct {
	gorm.Model
	Active bool `json:"active"`
}

// Use this function to place all your models
// in the database. This is useful for migrations.
func ImportModels() []any {
	return []any{
		&Beatmap{},
		&BeatmapSet{},
	}
}
