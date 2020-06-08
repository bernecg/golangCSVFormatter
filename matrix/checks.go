package matrix

import "strconv"

// VerifySquareMatrix Verify if Matrix has the same number of rows and
// columns, meaning is a square matrix
func VerifySquareMatrix(records [][]string) (isSquare bool) {
	if records != nil && len(records) == len(records[0]) {
		isSquare = true
	}
	return
}

// VerifyMatrixOnlyHasIntegers Verify if all values of the matrix are
// integers
func VerifyMatrixOnlyHasIntegers(records [][]string) (hasOnlyInt bool) {
	if records != nil {
		hasOnlyInt = true
		for _, row := range records {
			for _, value := range row {
				if _, err := strconv.Atoi(value); err != nil {
					hasOnlyInt = false
				}
			}
		}
	}
	return
}
