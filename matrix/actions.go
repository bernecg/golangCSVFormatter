package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Echo Return the matrix as a string in matrix format
func Echo(records [][]int) (response string) {
	if records != nil {
		for _, row := range records {
			var rowStr []string
			for _, value := range row {
				rowStr = append(rowStr, strconv.Itoa(value))
			}
			response = fmt.Sprintf("%s%s\n", response,
				strings.Join(rowStr, ","))
		}
	}
	return
}

// Invert Return the matrix as a string in a matrix format where the columns
// and rows are inverted
func Invert(records [][]int) (response string) {
	if records != nil {
		for indR, row := range records {
			var rowStr []string
			for indC := range row {
				rowStr = append(rowStr, strconv.Itoa(records[indC][indR]))
			}
			response = fmt.Sprintf("%s%s\n", response,
				strings.Join(rowStr, ","))
		}
	}
	return
}

// Flatten Return the matrix as a 1 line string, with values separated by
// commas
func Flatten(records [][]int) (response string) {
	if records != nil {
		var allValues []string
		for _, row := range records {
			for _, value := range row {
				allValues = append(allValues, strconv.Itoa(value))
			}
		}
		response = fmt.Sprintf("%s%s\n", response,
			strings.Join(allValues, ","))
	}
	return
}

// Sum Return the sum of the integers in the matrix
func Sum(records [][]int) (response string) {
	if records != nil {
		sum := 0
		for _, row := range records {
			for _, value := range row {
				sum += value
			}
		}
		response = fmt.Sprintf("%d\n", sum)
	}
	return
}

// Multiply Return the product of the integers in the matrix
func Multiply(records [][]int) (response string) {
	if records != nil {
		sum := 1
		for _, row := range records {
			for _, value := range row {
				sum *= value
			}
		}
		response = fmt.Sprintf("%d\n", sum)
	}
	return
}
