package utils

/*
Osu! API Token response
See https://osu.ppy.sh/docs/index.html#client-credentials-grant
*/
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

/*
BeatmapSet object
See https://osu.ppy.sh/docs/index.html#beatmapset-object
*/
type BeatmapSetResponse struct {
	Artist                string               `json:"artist"`
	ArtistUnicode         string               `json:"artist_unicode"`
	Covers                Covers               `json:"covers"`
	Creator               string               `json:"creator"`
	FavouriteCount        int                  `json:"favourite_count"`
	ID                    int                  `json:"id"`
	Nsfw                  bool                 `json:"nsfw"`
	Offset                int                  `json:"offset"`
	PlayCount             int                  `json:"play_count"`
	PreviewURL            string               `json:"preview_url"`
	Source                string               `json:"source"`
	Status                string               `json:"status"`
	Spotlight             bool                 `json:"spotlight"`
	Title                 string               `json:"title"`
	TitleUnicode          string               `json:"title_unicode"`
	UserID                int                  `json:"user_id"`
	Video                 bool                 `json:"video"`
	Beatmaps              []BeatmapResponse    `json:"beatmaps"`
	Converts              []string             `json:"converts"`
	CurrentNominations    []NominationResponse `json:"current_nominations"`
	CurrentUserAttributes string               `json:"current_user_attributes"`
	Description           string               `json:"description"`
	Discussions           string               `json:"discussions"`
	Events                string               `json:"events"`
	Genre                 string               `json:"genre"`
	HasFavourited         bool                 `json:"has_favourited"`
	Language              string               `json:"language"`
	Nominations           string               `json:"nominations"`
	PackTags              []string             `json:"pack_tags"`
	Ratings               string               `json:"ratings"`
	RecentFavourites      string               `json:"recent_favourites"`
	RelatedUsers          string               `json:"related_users"`
	User                  string               `json:"user"`
	TrackID               int                  `json:"track_id"`
}

/*
Beatmap object
See https://osu.ppy.sh/docs/index.html#beatmap
See https://osu.ppy.sh/docs/index.html#beatmapextended
*/
type BeatmapResponse struct {
	BeatmapSetID     int     `json:"beatmapset_id"`
	DifficultyRating float64 `json:"difficulty_rating"`
	ID               int     `json:"id"`
	Mode             string  `json:"mode"`
	Status           string  `json:"status"` // See Rank status for list of possible values.
	TotalLength      int     `json:"total_length"`
	UserID           int     `json:"user_id"`
	Version          string  `json:"version"`
	Accuracy         float64 `json:"accuracy"`
	AR               float64 `json:"ar"`
	BPM              float64 `json:"bpm"`
	Convert          bool    `json:"convert"`
	CountCircles     int     `json:"count_circles"`
	CountSliders     int     `json:"count_sliders"`
	CountSpinners    int     `json:"count_spinners"`
	CS               float64 `json:"cs"`
	DeletedAt        string  `json:"deleted_at"`
	Drain            float64 `json:"drain"`
	HitLength        int     `json:"hit_length"`
	IsScoreable      bool    `json:"is_scoreable"`
	LastUpdated      string  `json:"last_updated"`
	ModeInt          int     `json:"mode_int"`
	PassCount        int     `json:"passcount"`
	PlayCount        int     `json:"playcount"`
	Ranked           int     `json:"ranked"` // See Rank status for list of possible values.
	URL              string  `json:"url"`
}

/*
Cover on Beatmapset.

See https://osu.ppy.sh/docs/index.html#beatmapset-covers
*/
type Covers struct {
	Cover       string `json:"cover"`
	Cover2x     string `json:"cover@2x"`
	Card        string `json:"card"`
	Card2x      string `json:"card@2x"`
	List        string `json:"list"`
	List2x      string `json:"list@2x"`
	Slimcover   string `json:"slimcover"`
	Slimcover2x string `json:"slimcover@2x"`
}

/*
Nomination object

See https://osu.ppy.sh/docs/index.html#nomination
*/
type NominationResponse struct {
	beatmapset_id 	integer
	rulesets 	Ruleset[]
	reset 	boolean
	user_id 	integer
}

/*
Ruleset object
See https://osu.ppy.sh/docs/index.html#ruleset
*/
type Ruleset struct {
	Name string `json:"name"`
	Description string `json:"description"`
}