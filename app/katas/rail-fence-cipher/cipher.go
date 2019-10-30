package cipher

import (
	"fmt"
	"strings"
)

func Encode(input string, rows int) string {
	var result string
	padded := pad(input, rows)
	length := len(padded)
	middleRows := rows - 2
	maxColumn := length/(rows+middleRows) - 1
	resultSlice := strings.Split(padded, "")

	for i, r := range padded {
		var destination int
		row := i % (rows + middleRows)
		column := i / (2 + middleRows*2)
		letter := string(r)
		isFirstRow := row == 0
		isLastRow := row == rows-middleRows
		isRising := row >= rows

		if isRising {
			row = rows - row + 1
		}

		if isFirstRow {
			destination = column
		} else if isLastRow {
			destination = rows*(maxColumn+1) + column
		} else {
			destination = row*(maxColumn+1) + column*2

			if isRising {
				destination++
			}
		}

		resultSlice[destination] = letter
	}

	for _, s := range resultSlice {
		if s != "-" {
			result += s
		}
	}

	fmt.Println("maxColumn, resultSlice", maxColumn, resultSlice)

	return result
}

func Decode(input string, n int) string {
	// code
	return ""
}

func pad(input string, rows int) string {
	result := input
	length := len(input)
	middleRows := rows - 2
	lettersPerColumn := 2 + middleRows*2
	remainder := length % lettersPerColumn
	paddedLength := length

	if remainder > 0 {
		paddedLength = (length/lettersPerColumn + 1) * lettersPerColumn
	}

	for i := length; i < paddedLength; i++ {
		result += "-"
	}

	return result
}
