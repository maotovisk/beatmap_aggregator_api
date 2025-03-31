package model

type Beatmap struct {
	BaseModel
	BeatmapID     int      `json:"beatmap_id" gorm:"primaryKey"`
	BeatmapSetID  int      `json:"beatmapset_id"`
	Name          string   `json:"name"`
	Mappers       []Mapper `gorm:"many2many:beatmap_mappers;"`
	CS            float64  `json:"cs"`
	AR            float64  `json:"ar"`
	BPM           float64  `json:"bpm"`
	Length        int      `json:"length"`
	BackgroundURL string   `json:"background_url"`
}

type BeatmapSet struct {
	BaseModel
	BeatmapSetID int       `json:"beatmapset_id" gorm:"primaryKey"`
	Beatmaps     []Beatmap `json:"beatmaps" gorm:"foreignKey:BeatmapSetID"`
	Title        string    `json:"title"`
	Artist       string    `json:"artist"`
	Mapper       string    `json:"mapper"`
	Description  string    `json:"description"`
}

type Mapper struct {
	BaseModel
	UserID    int    `json:"user_id" gorm:"primaryKey"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}
