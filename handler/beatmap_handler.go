package handler

import (
	"net/http"
	"simple_api/database"
	"simple_api/model"
	"simple_api/utils"
)

func GetBeatmaps(w http.ResponseWriter, r *http.Request) {
	jh := utils.NewJsonHandler(w, r)
	db := database.GetDatabase()

	beatmaps := []model.Beatmap{}

	tx := db.Find(&beatmaps)
	if err := tx.Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jh.WriteResponse(beatmaps)
}

func InsertBeatmap(w http.ResponseWriter, r *http.Request) {
	type BodyParams struct {
		URL string `json:"url"`
	}

	jh := utils.NewJsonHandler(w, r)
	//db := database.GetDatabase()

	params := &BodyParams{}
	err := jh.ParseBody(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jh.WriteMessageWithStatus("Success! Parsed URL: "+params.URL, http.StatusOK)
}
