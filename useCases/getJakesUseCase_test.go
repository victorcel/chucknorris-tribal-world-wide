package useCases

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJacksUseCase(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`
					{"categories":[],
					"created_at":"2020-01-05 13:42:19.324003",
					"icon_url":"https://assets.chucknorris.host/img/avatar/chuck-norris.png",
					"id":"7hM0SAhiRNW8AE8UTlg6jw","updated_at":"2020-01-05 13:42:19.324003",
					"url":"https://api.chucknorris.io/jokes/7hM0SAhiRNW8AE8UTlg6jw",
					"value":"[Chuck Norris](poe://www.poe.com/_api/key_pck%20Norris.) 
					doesn't have a computer. He IS the computer."}`))
	}))

	defer mockServer.Close()

	jokes, err := GetJacksUseCase(mockServer.URL)
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}

	if len(jokes) == 0 {
		t.Error("No jokes returned")
	}
}

func TestGetJokes_ErrorMakingRequest(t *testing.T) {
	ctx := context.Background()
	url := "http://localhost:12345"

	_, err := getJokes(ctx, url)

	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}
