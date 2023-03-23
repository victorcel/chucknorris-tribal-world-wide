package useCases

import (
	"context"
	"encoding/json"
	"github.com/victorcel/chucknorris/models"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func getJokes(ctx context.Context, url string) (models.JakeModel, error) {
	client := &http.Client{}

	modelJake := models.JakeModel{}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println(err.Error())
		return modelJake, err
	}

	response, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return modelJake, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(response.Body)

	errJson := json.NewDecoder(response.Body).Decode(&modelJake)
	if err != nil {
		return modelJake, errJson
	}

	return modelJake, nil
}

func GetJacksUseCase(url string) ([]models.ResponseJake, error) {
	var jokes []models.ResponseJake
	const NumberObject = 25
	jokeIDs := make(map[string]bool)
	jokeChan := make(chan models.JakeModel)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(NumberObject)

	for i := 0; i < NumberObject; i++ {
		go func() {
			defer wg.Done()
			joke, err := getJokes(ctx, url)
			if err != nil {
				log.Println(err.Error())
				return
			}
			jokeChan <- joke
		}()
	}

	go func() {
		wg.Wait()
		close(jokeChan)
	}()

	for joke := range jokeChan {
		if !jokeIDs[joke.Id] {
			jokeIDs[joke.Id] = true
			jokes = append(jokes, models.ResponseJake{
				Id:    joke.Id,
				Url:   joke.Url,
				Value: joke.Value,
			})
		}
	}

	return jokes, nil
}
