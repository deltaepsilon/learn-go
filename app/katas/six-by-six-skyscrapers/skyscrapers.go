package skyscrapers

import "fmt"

type eliminationMap [6]vector
type vector [6]cell
type cell struct {
	row    int
	column int
	values [6]bool
}

func SolvePuzzle(clues []int) [][]int {
	var result [][]int
	eMap := getEliminationMap()

	for i, clue := range clues {
		switch clue {
		case 6:
			fmt.Println("sixes")
			eMap = sixes(eMap, clue, i)
			logMap(eMap)

		case 5:
			fmt.Println("fives")
			eMap = fives(eMap, clue, i)
			logMap(eMap)

		case 4:
			fmt.Println("fours")
			eMap = fours(eMap, clue, i)
			logMap(eMap)

		case 3:
			fmt.Println("threes")
			eMap = threes(eMap, clue, i)
			logMap(eMap)

		case 2:
			fmt.Println("twos")
			eMap = twos(eMap, clue, i)
			logMap(eMap)

		case 1:
			fmt.Println("ones")
			eMap = ones(eMap, clue, i)
			logMap(eMap)

		}
	}

	fmt.Println("solved")
	logMap(eMap)

	// Return 6x6 matrix here ...
	return result
}

func sixes(eMap eliminationMap, clue, clueNumber int) eliminationMap {
	return processClue(eMap, clue, clueNumber, func(v vector, clue int) vector {
		for i := range v {
			var values [6]bool

			for j := 0; j < 6; j++ {
				if i == j {
					values[j] = true
				} else {
					values[j] = false
				}
			}

			v[i].values = values
		}

		return v
	})
}

func fives(eMap eliminationMap, clue, clueNumber int) eliminationMap {
	return processClue(eMap, clue, clueNumber, func(v vector, clue int) vector {
		v[0].values[5] = false
		v[1].values[5] = false
		v[2].values[5] = false
		v[3].values[5] = false

		return v
	})
}

func fours(eMap eliminationMap, clue, clueNumber int) eliminationMap {
	return processClue(eMap, clue, clueNumber, func(v vector, clue int) vector {
		v[0].values[5] = false
		v[1].values[5] = false
		v[2].values[5] = false

		return v
	})
}

func threes(eMap eliminationMap, clue, clueNumber int) eliminationMap {
	return processClue(eMap, clue, clueNumber, func(v vector, clue int) vector {
		v[0].values[5] = false
		v[1].values[5] = false

		return v
	})
}

func twos(eMap eliminationMap, clue, clueNumber int) eliminationMap {
	return processClue(eMap, clue, clueNumber, func(v vector, clue int) vector {
		v[0].values[5] = false

		return v
	})
}

func ones(eMap eliminationMap, clue, clueNumber int) eliminationMap {
	return processClue(eMap, clue, clueNumber, func(v vector, clue int) vector {
		v[0].values = [6]bool{false, false, false, false, false, true}

		return v
	})
}

func processClue(eMap eliminationMap, clue, clueNumber int, processVector func(vector, int) vector) eliminationMap {
	var v vector
	var processed bool
	clueBlock := clueNumber / 6
	vectorNumber := clueNumber % 6

	switch clueBlock {
	case 0:
		v = getColumn(eMap, vectorNumber)
		v = processVector(v, clue)
		processed = true

	case 1:
		v = getReversedRow(eMap, vectorNumber)
		v = processVector(v, clue)
		processed = true

	case 2:
		v = getReversedColumn(eMap, 5-vectorNumber)
		v = processVector(v, clue)
		processed = true

	case 3:
		v = getRow(eMap, 5-vectorNumber)
		v = processVector(v, clue)
		processed = true

	}

	if processed {
		eMap = replaceVector(eMap, v)
	}

	fmt.Println("")
	fmt.Println("clueBlock", clueBlock)
	fmt.Println("vectorNumber", vectorNumber)
	eMap = eliminate(eMap)

	return eMap
}

func eliminate(eMap eliminationMap) eliminationMap {
	result := solveMap(eMap)

	for rowNumber := 0; rowNumber < 6; rowNumber++ {
		for columnNumber := 0; columnNumber < 6; columnNumber++ {
			result = eliminateSolvedCell(result, rowNumber, columnNumber)
		}
	}

	if hasChanged(eMap, result) {
		result = eliminate(result)
	}

	return result
}

func solveMap(eMap eliminationMap) eliminationMap {
	for i := 0; i < 6; i++ {
		row := getRow(eMap, i)
		eMap = replaceVector(eMap, solveVector(row))

		column := getColumn(eMap, i)
		eMap = replaceVector(eMap, solveVector(column))
	}

	return eMap
}

func solveVector(v vector) vector {
	var frequencies [6]int

	for _, cell := range v {
		for i, value := range cell.values {
			if value {
				frequencies[i]++
			}
		}
	}

	for i, frequency := range frequencies {
		if frequency == 1 {
			for j := 0; j < 6; j++ {
				cell := v[j]

				if cell.values[i] == true {
					for k := 0; k < 6; k++ {
						if i != k {
							v[j].values[k] = false
						}

					}

				}
			}
		}
	}

	return v
}

func replaceVector(eMap eliminationMap, v vector) eliminationMap {
	for _, cell := range v {
		eMap[cell.row][cell.column] = cell
	}

	return eMap
}

func eliminateSolvedCell(eMap eliminationMap, rowNumber, columnNumber int) eliminationMap {
	c := eMap[rowNumber][columnNumber]
	solution := getSolution(c)

	if solution != -1 {
		eMap = processEliminationMap(eMap, func(c cell) cell {
			if c.row == rowNumber && c.column != columnNumber {
				c.values[solution] = false
			}

			if c.column == columnNumber && c.row != rowNumber {
				c.values[solution] = false
			}

			return c
		})
	}

	return eMap
}

func getSolution(c cell) int {
	var result int = -1
	truthy := 0

	for i, value := range c.values {
		if value {
			truthy++
			result = i
		}
	}

	if truthy != 1 {
		result = -1
	}

	return result
}

func processEliminationMap(eMap eliminationMap, processCell func(c cell) cell) eliminationMap {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			eMap[i][j] = processCell(eMap[i][j])
		}
	}

	return eMap
}

func hasChanged(a, b eliminationMap) bool {
	return fmt.Sprint(a) != fmt.Sprint(b)
}

func getColumn(eMap eliminationMap, columnNumber int) vector {
	var v vector

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if j == columnNumber {
				v[i] = eMap[i][j]
			}
		}
	}

	return v
}

func getReversedColumn(eMap eliminationMap, columnNumber int) vector {
	var v vector

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if j == columnNumber {
				v[5-i] = eMap[i][j]
			}
		}
	}

	return v
}

func getRow(eMap eliminationMap, rowNumber int) vector {
	var v vector

	for i, cell := range eMap[rowNumber] {
		v[i] = cell
	}

	return v
}

func getReversedRow(eMap eliminationMap, rowNumber int) vector {
	var v vector

	for i, cell := range eMap[rowNumber] {
		v[5-i] = cell
	}

	return v
}

// Logging

func getEliminationMap() eliminationMap {
	var result eliminationMap

	for i, row := range result {
		for j := range row {
			result[i][j] = cell{
				row: i, column: j, values: [6]bool{true, true, true, true, true, true}}
		}
	}

	return result
}

func logMap(eMap eliminationMap) {
	fmt.Println("")

	for _, row := range eMap {
		logVector(row)
		fmt.Println("")
	}

	fmt.Println("")
}

func logVector(v vector) {
	for _, cell := range v {
		logCell(cell)
	}
}

func logCell(c cell) {
	fmt.Printf("[%vx%v ", c.row, c.column)
	for j, value := range c.values {

		if value == true {
			fmt.Printf("%v", j+1)
		} else {
			fmt.Printf(" ")

		}
	}
	fmt.Printf("]")
}
