package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"simple_api/config"
)

var token string

func getToken() string {
	config := config.GetConfig()

	client_id := config.Osu.ClientID
	client_secret := config.Osu.ClientSecret
	grant_type := "client_credentials"
	scope := "public"

	client := &http.Client{}

	requestBody := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=%s&scope=%s", client_id, client_secret, grant_type, scope)

	req, err := http.NewRequest("POST", "https://osu.ppy.sh/oauth/token", bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: Code: %s", resp.Status)
		log.Printf("error: Body: %s", string(body))
	}

	return tokenResponse.AccessToken
}

func postRequest[T any](url string, requestBody any) (T, error) {
	var bodyJson []byte
	if requestBody != nil {
		var err error
		bodyJson, err = json.Marshal(requestBody)
		if err != nil {
			log.Fatal(err)
		}
	}

	bodyBuffer := bytes.NewBuffer(bodyJson)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bodyBuffer)
	if err != nil {
		log.Fatal(err)
	}

	if token == "" {
		token = getToken()
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var result T
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: Code: %s", resp.Status)
		log.Printf("error: Body: %s", io.ReadCloser(resp.Body))
	}
	return result, nil
}

func getRequest[T any](url string, params map[string]string) (T, error) {

	if params != nil || len(params) > 0 {
		url += "?"
		for key, value := range params {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	if token == "" {
		token = getToken()
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var result T
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: Code: %s", resp.Status)
		log.Printf("error: Body: %s", io.ReadCloser(resp.Body))
	}
	return result, nil
}

func GetOsuBeatmap(beatmapID int) (BeatmapResponse, error) {
	url := fmt.Sprintf("https://osu.ppy.sh/api/v2/beatmaps/%d", beatmapID)
	params := map[string]string{
		"b": fmt.Sprintf("%d", beatmapID),
	}

	return getRequest[BeatmapResponse](url, params)
}

func GetOsuBeatmapSets(beatmapSetID int) (BeatmapSetResponse, error) {
	url := fmt.Sprintf("https://osu.ppy.sh/api/v2/beatmapsets/%d", beatmapSetID)
	params := map[string]string{
		"b": fmt.Sprintf("%d", beatmapSetID),
	}

	return getRequest[BeatmapSetResponse](url, params)
}
