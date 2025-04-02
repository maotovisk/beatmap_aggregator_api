package utils_test

import (
	"beatmap_aggregator_api/utils"
	"testing"
)

func TestExtractBeatmapSetIDFromURL(t *testing.T) {
	id, err := utils.ExtractBeatmapSetIDFromURL("https://osu.ppy.sh/beatmapsets/123456#osu/12345")
	if err != nil || id != "123456" {
		t.Errorf("expected %s, got %s", "123456", id)
	}
}
