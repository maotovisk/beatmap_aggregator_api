package utils

import (
	"beatmap_aggregator_api/config"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/go-resty/resty/v2"
)

var token string

func getToken() string {
	config := config.GetConfig()
	client := resty.New()

	var tokenResponse TokenResponse

	res, err := client.R().
		SetFormData(map[string]string{
			"client_id":     config.Osu.ClientID,
			"client_secret": config.Osu.ClientSecret,
			"grant_type":    "client_credentials",
			"scope":         "public",
		}).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Accept", "application/json").
		SetResult(&tokenResponse).
		Post("https://osu.ppy.sh/oauth/token")
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode() != http.StatusOK {
		log.Fatalf("error: Code: %s", res.Status())
		log.Printf("error: Body: %s", res.Body())
	}

	return tokenResponse.AccessToken
}

func postRequest[T any](url string, requestBody any) (T, error) {
	var response T
	client := resty.New()

	if token == "" {
		token = getToken()
	}

	res, err := client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		SetResult(&response).
		Post(url)
	if err != nil {
		return response, err
	}

	if res.StatusCode() != http.StatusOK {
		log.Printf("error: Code: %s", res.Status())
		log.Printf("error: Body: %s", res.Body())
		return response, fmt.Errorf("error: Code: %s", res.Status())
	}

	return response, nil
}

func getRequest[T any](url string, params map[string]string) (T, error) {
	var response T
	client := resty.New()

	if token == "" {
		token = getToken()
	}

	res, err := client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("Accept", "application/json").
		SetQueryParams(params).
		SetResult(&response).
		Get(url)
	if err != nil {
		return response, err
	}

	if res.StatusCode() != http.StatusOK {
		log.Printf("error: Code: %s", res.Status())
		log.Printf("error: Body: %s", res.Body())
		return response, fmt.Errorf("error: Code: %s", res.Status())
	}

	return response, nil
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

func ExtractBeatmapSetIDFromURL(url string) (string, error) {
	// Example URL: https://osu.ppy.sh/beatmapsets/123456#osu/123456
	beatmapRegex := regexp.MustCompile(`https://osu.ppy.sh/beatmapsets/(\d*)#.*/.*`)

	matches := beatmapRegex.FindStringSubmatch(url)
	if len(matches) != 2 {
		return "", fmt.Errorf("couldn't parse the beatmap URL")
	}

	return matches[1], nil
}
