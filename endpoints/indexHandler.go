package endpoints

import (
	"encoding/json"
	"golang-challenge/models"
	"net/http"
)

// IndexHandler is the first endpoint to the Challenger API.
// This method is responsible for display the result of printNumbers according to the challenger requested.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(printNumbers())
}

// printNumbers is responsible for generate numbers 1 to 100. It returns a models.Exports type
func printNumbers() models.Exports {
	rangeN := []interface{}{}

	for i := 1; i <= 100; i++ {
		rangeN = append(rangeN, i)
	}

	dataToExport := models.Exports{
		models.ExportData{
			Data:       rangeN,
			StatusCode: 200,
			Message:    `The default parameter-less behaviour is to print the numbers from 1 to 100.`,
		},
	}
	return dataToExport
}
