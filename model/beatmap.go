package model

type Beatmap struct {
	BaseModel
	BeatmapID    int    `json:"beatmap_id" gorm:"primaryKey"`
	BeatmapSetID int    `json:"beatmapset_id"`
	Title        string `json:"title"`
	Artist       string `json:"artist"`
	Mapper       string `json:"mapper"`
	Tags         string `json:"tags"`
}

type BeatmapSet struct {
	BaseModel
	BeatmapSetID int       `json:"beatmapset_id" gorm:"primaryKey"`
	Beatmaps     []Beatmap `json:"beatmaps" gorm:"foreignKey:BeatmapSetID"`
}
