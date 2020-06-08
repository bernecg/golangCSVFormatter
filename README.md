# golangCSVFormatter

A web server that accepts a CSV file with a square matrix and do some actions on it, choosen by the endpoint.

## How it works

Given an uploaded csv file
```
1,2,3
4,5,6
7,8,9
```

__Echo__: Return the matrix as a string in matrix format.
```
// Expected output
1,2,3
4,5,6
7,8,9
```

__Invert__: Return the matrix as a string in matrix format where the columns and rows are inverted
```
// Expected output
1,4,7
2,5,8
3,6,9
```

__Flatten__: Return the matrix as a 1 line string, with values separated by commas.
```
// Expected output
1,2,3,4,5,6,7,8,9
```

__Sum__: Return the sum of the integers in the matrix
```
// Expected output
45
```

__Multiply__: Return the product of the integers in the matrix
```
// Expected output
362880
```

The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. _matrix.csv_ is example valid input.

## How to run

Run web server
```
go run .
```

Send request
```
curl -F 'file=@matrix.csv' "localhost:8080/echo"
```

## How to run tests

```
go test ./test -v
```

## TODOs / Ideas
- Create more Unit Tests
- Maybe separated better functionalities (e.g: file things)
- Use channels and goroutines on some cases, especially on arithmetics ones.
- Create a Makefile file
- Better commentaries and documentation
