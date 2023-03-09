package controller

import (
	"encoding/json"
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
	"net/http"
	"strconv"
)

func GetRankingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// get query parameters
	queryParams := r.URL.Query()
	numberOfData, err := strconv.Atoi(queryParams.Get("number"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	// TODO: Add query validation here

	// extract keyboard data from database based on their net ranking
	var keyboards []model.Keyboard
	keyboards = model.GetRanks(numberOfData)
	if keyboards == nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("No keyboard found"))
		return
	}

	// Return the slice of keyboards as a JSON response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(keyboards)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error encoding response: " + err.Error()))
		return
	}
}
