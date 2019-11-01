package cipher

import (
	"math"
	"strings"
)

func Encode(input string, rows int) string {
	length := len(input)
	lettersPerColumn := getLettersPerColumn(rows)
	result := strings.Split(input, "")
	rowCounts := getRowCounts(input, rows)

	for i, r := range input {
		var destination int
		row, isRising := getEncodeRow(i, rows)
		column := getEncodeColumn(i, rows)
		letter := string(r)
		isFirstRow := row == 0
		isLastRow := row == lettersPerColumn/2

		if isFirstRow {
			destination = column

		} else if isLastRow {
			rowCount := rowCounts[row]
			destination = length - rowCount + column

		} else {
			countInPrecedingRows := getCountInPrecedingRows(input, i, rows)
			destination = countInPrecedingRows + column

			if column > 0 {
				destination += column
			}

			if isRising {
				destination++
			}
		}

		result[destination] = letter
	}

	return strings.Join(result, "")
}

func Decode(input string, rows int) string {
	lettersPerColumn := getLettersPerColumn(rows)
	result := strings.Split(input, "")

	for i, r := range input {
		var destination int
		row := getDecodeRow(input, i, rows)
		column, isRising := getDecodeColumn(input, i, rows)
		letter := string(r)
		isFirstRow := row == 0
		isLastRow := row == rows-1

		if isFirstRow {
			destination = column * lettersPerColumn

		} else if isLastRow {
			destination = lettersPerColumn/2 + column*lettersPerColumn

		} else {
			if isRising {
				intraColumnOffset := (rows-row)*2 - 2

				destination = column*lettersPerColumn + row + intraColumnOffset
			} else {
				destination = column*lettersPerColumn + row
			}

		}

		result[destination] = letter

	}

	return strings.Join(result, "")
}

func getDecodeColumn(input string, i, rows int) (int, bool) {
	var result int
	var isRising bool

	row := getDecodeRow(input, i, rows)
	isFirstRow := row == 0
	isLastRow := row == rows-1

	if isFirstRow {
		result = i
	} else if isLastRow {
		rowStart := getDecodeRowStart(input, i, rows)
		result = int(math.Max(float64(i-rowStart), 0.0))

	} else {
		rowStart := getDecodeRowStart(input, i, rows)
		startIsOdd := rowStart%2 == 1
		isOdd := i%2 == 1
		isRising = isOdd != startIsOdd
		value := i

		if isRising {
			value--
		}

		result = (value - rowStart) / 2
	}

	return result, isRising
}

func getDecodeRowStart(input string, i, rows int) int {
	var rowStart int
	row := getDecodeRow(input, i, rows)
	rowCounts := getRowCounts(input, rows)

	for i, count := range rowCounts {
		if i == row {
			break
		} else {
			rowStart += count
		}
	}

	return rowStart
}

func getDecodeRow(input string, i, rows int) int {
	var result int = rows - 1
	var accumulator int

	rowCounts := getRowCounts(input, rows)

	for row, count := range rowCounts {
		accumulator += count

		if i < accumulator {
			result = row

			break
		}
	}

	return result
}

func getCountInPrecedingRows(input string, i, rows int) int {
	var result int
	row, _ := getEncodeRow(i, rows)

	if row > 0 {
		rowCounts := getRowCounts(input, rows)

		for i := 0; i < row; i++ {
			result += rowCounts[i]
		}
	}

	return result
}

func getRowCounts(input string, rows int) []int {
	var result []int

	for i := 0; i < len(input); i++ {
		row, _ := getEncodeRow(i, rows)

		if len(result) <= row {
			result = append(result, 0)
		}

		result[row]++
	}

	return result
}

func getColumnCounts(input string, rows int) []int {
	var result []int

	for i := 0; i < len(input); i++ {
		column := getEncodeColumn(i, rows)

		if len(result) <= column {
			result = append(result, 0)
		}

		result[column]++
	}

	return result
}

func getEncodeRow(i, rows int) (int, bool) {
	lettersPerColumn := getLettersPerColumn(rows)
	row := i % lettersPerColumn
	isRising := row >= rows

	if isRising {
		row = lettersPerColumn - row
	}

	return row, isRising
}

func getEncodeColumn(i, rows int) int {
	lettersPerColumn := getLettersPerColumn(rows)

	return i / lettersPerColumn
}

func getLettersPerColumn(rows int) int {
	return 2*rows - 2
}
