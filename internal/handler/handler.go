package handler

import (
	"sort"
	"net/http"
	"encoding/json"
	"alexdenkk/test-task/internal/repository"
	"alexdenkk/test-task/internal/model"
)

type Handler struct {
	Repository *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{
		Repository: repo,
	}
}

type AddNumberRequest struct {
	Value int `json:"value"`
}

func (handler *Handler) sortAndExtractNumbers(numbers []model.Number) []int {
	if len(numbers) == 0 {
		return []int{}
	}
	
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i].Value < numbers[j].Value
	})

	var end []int

	for _, num := range numbers {
		end = append(end, num.Value)
	}

	return end
}

func (handler *Handler) AddNumber(w http.ResponseWriter, r *http.Request) {
	var data AddNumberRequest

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.Repository.CreateNumber(
		r.Context(),
		model.Number{
			Value: data.Value,
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	numbers, err := handler.Repository.GetAllNumbers(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	extracted := handler.sortAndExtractNumbers(numbers)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(extracted)
}
