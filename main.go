package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"

	matrix "./matrix"
)

// convertRecordsToInt Convert the string Matrix to an integer Matrix
func convertRecordsToInt(w http.ResponseWriter, records [][]string) (recordsInt [][]int) {
	if !matrix.VerifySquareMatrix(records) {
		w.Write([]byte(fmt.Sprintf("error %s\n", "matrix is not a square")))
		return nil
	} else if !matrix.VerifyMatrixOnlyHasIntegers(records) {
		w.Write([]byte(fmt.Sprintf("error %s\n",
			"matrix has some values that are not integers")))
		return nil
	}

	for _, row := range records {
		var rowInt []int
		for _, value := range row {
			valueInt, err := strconv.Atoi(value)
			if err == nil {
				rowInt = append(rowInt, valueInt)
			} else {
				w.Write([]byte(fmt.Sprintf("error %s\n", err.Error())))
				return nil
			}
		}
		recordsInt = append(recordsInt, rowInt)
	}
	return
}

// getRecordsFromFile Return the matrix in CSV file
func getRecordsFromFile(w http.ResponseWriter, r *http.Request) (records [][]int) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s\n", err.Error())))
		return nil
	}
	defer file.Close()

	var recordsStr [][]string
	recordsStr, err = csv.NewReader(file).ReadAll()
	records = convertRecordsToInt(w, recordsStr)

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s\n", err.Error())))
		return nil
	}
	return
}

// newHandler Create a new Handler for given endpoint and the action for it
func newHandler(endpoint string, action func([][]int) string) {
	http.HandleFunc("/"+endpoint, func(w http.ResponseWriter, r *http.Request) {
		var response string
		if records := getRecordsFromFile(w, r); records != nil {
			response = action(records)
		}
		fmt.Fprint(w, response)
	})
}

func main() {
	newHandler("echo", matrix.Echo)
	newHandler("invert", matrix.Invert)
	newHandler("flatten", matrix.Flatten)
	newHandler("sum", matrix.Sum)
	newHandler("multiply", matrix.Multiply)
	http.ListenAndServe(":8080", nil)
}
