package useCases

import (
	"encoding/json"
	"fmt"
	"github.com/victorcel/chucknorris/models"
	"io"
	"net/http"
	"os"
)

func getJokes() (models.JakeModel, error) {

	url := os.Getenv("URL_API")

	client := &http.Client{}

	modelJake := models.JakeModel{}

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Println(err.Error())
		return modelJake, err
	}

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err.Error())
		return modelJake, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(response.Body)

	errJson := json.NewDecoder(response.Body).Decode(&modelJake)
	if err != nil {
		return modelJake, errJson
	}

	return modelJake, nil
}

func GetJacksUseCase() ([]models.ResponseJake, error) {
	var jokes []models.ResponseJake
	const NumberObject = 25
	jokeIDs := make(map[string]bool)
	jokeChan := make(chan models.JakeModel)

	for i := 0; i < NumberObject; i++ {
		go func() {
			joke, err := getJokes()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			jokeChan <- joke
		}()
	}

	for i := 0; i < NumberObject; i++ {
		joke := <-jokeChan
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
