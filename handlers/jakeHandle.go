package handlers

import (
	"encoding/json"
	"github.com/victorcel/chucknorris/useCases"
	"net/http"
	"os"
)

func GetJacksHandle() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		writer.Header().Set("Content-Type", "application/json")
		jokes, err := useCases.GetJacksUseCase(os.Getenv("URL_API"))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(writer).Encode(jokes)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
