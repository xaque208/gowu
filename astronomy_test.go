package gowu

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestJsonDecode(t *testing.T) {
	astronomyJson := LoadFixtureFile("astronomy.json")
	// fmt.Println(string(astronomyJson))

	var f AstronomyResponse
	err := json.Unmarshal(astronomyJson, &f)
	if err != nil {
		log.Error(err)
	}

	fmt.Printf("%+v", f)

	if f.Response.Features.Astronomy != 1 {
		t.Errorf("Expected %d, got %d", 1, f.Response.Features.Astronomy)
	}

	if f.SunPhase.Sunrise.Hour != "5" {
		t.Errorf("Expected %s, got %s", "5", f.SunPhase.Sunrise.Hour)
	}

	if f.SunPhase.Sunrise.Minute != "47" {
		t.Errorf("Expected %s, got %s", "47", f.SunPhase.Sunrise.Hour)
	}

}
