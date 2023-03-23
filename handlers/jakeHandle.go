package handlers

import (
	"encoding/json"
	"github.com/victorcel/chucknorris/models"
	"github.com/victorcel/chucknorris/useCases"
	"net/http"
)

var responseJacks []models.ResponseJake

func GetJacksHandle() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		writer.Header().Set("Content-Type", "application/json")
		jokes, err := useCases.GetJacksUseCase()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(jokes)
	}
}
