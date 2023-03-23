package models

import (
	"encoding/json"
	"testing"
)

func TestJakeModel(t *testing.T) {
	jsonStr := `{
        "categories": [],
        "created_at": "2022-03-22 10:00:00",
        "icon_url": "http://example.com/icon.png",
        "id": "1234",
        "updated_at": "2022-03-22 11:00:00",
        "url": "http://example.com/joke/1234",
        "value": "This is a joke"
    }`

	// Unmarshal the JSON string into a JakeModel struct
	var jakeModel JakeModel
	err := json.Unmarshal([]byte(jsonStr), &jakeModel)

	if err != nil {
		t.Errorf("Expected err to be nil, but got: %v", err)
	}

	if len(jakeModel.Categories) != 0 {
		t.Errorf("Expected jakeModel.Categories to be empty, but got: %v", jakeModel.Categories)
	}

	if jakeModel.CreatedAt != "2022-03-22 10:00:00" {
		t.Errorf("Expected jakeModel.CreatedAt to be %s, but got: %s", "2022-03-22 10:00:00", jakeModel.CreatedAt)
	}

	if jakeModel.IconUrl != "http://example.com/icon.png" {
		t.Errorf("Expected jakeModel.IconUrl to be %s, but got: %s", "http://example.com/icon.png", jakeModel.IconUrl)
	}

	if jakeModel.Id != "1234" {
		t.Errorf("Expected jakeModel.Id to be %s, but got: %s", "1234", jakeModel.Id)
	}

	if jakeModel.UpdatedAt != "2022-03-22 11:00:00" {
		t.Errorf("Expected jakeModel.UpdatedAt to be %s, but got: %s", "2022-03-22 11:00:00", jakeModel.UpdatedAt)
	}

	if jakeModel.Url != "http://example.com/joke/1234" {
		t.Errorf("Expected jakeModel.Url to be %s, but got: %s", "http://example.com/joke/1234", jakeModel.Url)
	}

	if jakeModel.Value != "This is a joke" {
		t.Errorf("Expected jakeModel.Value to be %s, but got: %s", "This is a joke", jakeModel.Value)
	}
}
