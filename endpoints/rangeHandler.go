package endpoints

import (
	"encoding/json"
	"golang-challenge/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// RangeHandler is the endpoint to the Challenger API. This method is calling when params to and from are displayed.
// This method is responsible for display the result of printNumbersUponRange according to the challenger requested.
func RangeHandler(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(printNumbersUponRange(params))
}

// printNumbersUponRange is responsible for generate the rule requested.
// The rule is: for multiples of 3, print the word 'Pé' instead of the number. For multiples of 5, print the word 'Do' instead of the number. For multiples of both 3 and 5, print 'PéDo'. It returns a models.Exports type
// Only allowed positive numbers in the params
func printNumbersUponRange(param map[string]string) models.Exports {
	rangeN := []interface{}{}

	from, _ := strconv.Atoi(param["fromRange"])
	to, _ := strconv.Atoi(param["toRange"])

	if isPositiveNumber(from) == false || isPositiveNumber(to) == false {
		dataToExport := models.Exports{
			models.ExportData{
				Data:       rangeN,
				StatusCode: 422,
				Message:    `Only allowed positive numbers in the params`,
			},
		}
		return dataToExport
	}

	for i := from; i <= to; i++ {
		if i%15 == 0 {
			rangeN = append(rangeN, "PéDo")
		} else if i%3 == 0 {
			rangeN = append(rangeN, "Pé")
		} else if i%5 == 0 {
			rangeN = append(rangeN, "Do")
		} else {
			rangeN = append(rangeN, i)
		}
	}
	rangeN = append(rangeN)

	dataToExport := models.Exports{
		models.ExportData{
			Data:       rangeN,
			StatusCode: 200,
			Message:    `For multiples of 3, print the word 'Pé' instead of the number. For multiples of 5, print the word 'Do' instead of the number. For multiples of both 3 and 5, print 'PéDo'`,
		},
	}
	return dataToExport
}

//It methods check if param requested is a positive or negative number. It returns a bool type
func isPositiveNumber(number int) bool {
	if number < 0 {
		return false
	}
	return true
}
