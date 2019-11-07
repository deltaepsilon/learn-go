package skyscrapers

import (
	"fmt"
	"math"
)

type eliminationMap [6]vector
type vector [6]cell
type cell struct {
	row    int
	column int
	values [6]bool
}

func SolvePuzzle(clues []int) [][]int {
	eMap := getEliminationMap()
	eMap = processClues(eMap, clues)

	result, isSolved := flattenMap(eMap)

	if isSolved == false {

		fmt.Println("isSolved", isSolved)
		fmt.Println("result", result)
		logMap(eMap, clues)
	}

	return result
}

func processClues(eMap eliminationMap, clues []int) eliminationMap {
	initialMap := eMap

	for i, clue := range clues {
		switch clue {
		case 6:
			// fmt.Println("sixes")
			eMap = sixes(eMap, clue, i)
			// logMap(eMap, clues)

		case 5:
			eMap = fives(eMap, clue, i)
			// fmt.Println("fives")
			// logMap(eMap, clues)

		case 4:
			// fmt.Println("fours")
			eMap = fours(eMap, clue, i)
			// logMap(eMap, clues)

		case 3:
			// fmt.Println("threes")
			eMap = threes(eMap, clue, i)
			// logMap(eMap, clues)

		case 2:
			// fmt.Println("twos")
			eMap = twos(eMap, clue, i)
			// logMap(eMap, clues)

		case 1:
			// fmt.Println("ones")
			eMap = ones(eMap, clue, i)
			// logMap(eMap, clues)

		}
	}

	if hasChanged(eMap, initialMap) {
		eMap = processClues(eMap, clues)
	}

	return eMap
}

func flattenMap(eMap eliminationMap) ([][]int, bool) {
	var isSolved bool = true
	result := make([][]int, 6)

Solve:
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			s := getSolution(eMap[i][j])

			if s == -1 {
				isSolved = false

				break Solve
			} else {
				if len(result) < i+1 {
					result = append(result, make([]int, 6))
				}

				result[i] = append(result[i], s+1)
			}
		}
	}

	return result, isSolved
}

func sixes(eMap eliminationMap, clue, clueNumber int) eliminationMap {
	return processClue(eMap, clue, clueNumber, func(v vector, clue int) vector {
		return enforceRising(v, 5)
	})
}

func fives(eMap eliminationMap, clue, clueNumber int) eliminationMap {
	return processClue(eMap, clue, clueNumber, func(v vector, clue int) vector {
		v[0].values[5] = false
		v[1].values[5] = false
		v[2].values[5] = false
		v[3].values[5] = false

		for i := 0; i < 6; i++ {
			s := getSolution(v[i])

			if s == 5 && i == 5 {
				v = enforceRising(v, 3)
			}

			if s == 5 && i == 4 {
				v = enforceRising(v, i)
			}
		}

		return v
	})
}

func fours(eMap eliminationMap, clue, clueNumber int) eliminationMap {
	return processClue(eMap, clue, clueNumber, func(v vector, clue int) vector {
		v[0].values[3] = false
		v[0].values[4] = false
		v[0].values[5] = false
		v[1].values[4] = false
		v[1].values[5] = false
		v[2].values[5] = false

		for i := 2; i < 4; i++ {
			if v[i].values[4] == false {
				v[i+1].values[5] = false
			}
		}

		for i := 0; i < 6; i++ {
			s := getSolution(v[i])

			if s == 5 && i == 3 {
				v = enforceRising(v, i)

			}
		}

		return v
	})
}

func threes(eMap eliminationMap, clue, clueNumber int) eliminationMap {
	return processClue(eMap, clue, clueNumber, func(v vector, clue int) vector {
		v[0].values[5] = false
		v[1].values[5] = false

		for i := 0; i < 6; i++ {
			s := getSolution(v[i])

			if s == 5 && i == 2 {
				v = enforceRising(v, i)
			}
		}

		return v
	})
}

func twos(eMap eliminationMap, clue, clueNumber int) eliminationMap {
	return processClue(eMap, clue, clueNumber, func(v vector, clue int) vector {
		v[0].values[5] = false

		var hasSolvedSix bool
		var firstCellMax int
		var maxInternalSolution int

		for i := 5; i > 0; i-- {
			s := getSolution(v[i])

			if hasSolvedSix {
				v[i] = trimCellMax(v[i], firstCellMax)
			}

			if hasSolvedSix && s != -1 {
				maxInternalSolution = int(math.Max(float64(s), float64(maxInternalSolution)))
			}

			if s == 5 {
				hasSolvedSix = true
				firstCellMax = getCellMax(v[0])
			}
		}

		for i := 1; i < 6; i++ { //
			if v[i].values[5] == true {
				break
			} else {
				// TODO
				// Count the cells between the first and the first possible 6
				// trim the first cell from 1 to n based on that number
				// For example, if there are two cells in between the first
				// cell and the first possible six, then the first cell must be
				// at least a three

				// TODO consider checking for missing possibilities in the interstitial
				// cells.
				// That could potentially require the first cell min to be higher

				// TODO all cells in between the first and first possible six
				// must be less than the firstCellMax

				fmt.Println("firstCellMax", firstCellMax)
				logCell(v[i])
				fmt.Println("")

				// v[i] = trimCellMax(v[i], firstCellMax)
			}
		}

		if maxInternalSolution > 0 {
			v[0] = trimCellMin(v[0], maxInternalSolution)
		}

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

func enforceRising(v vector, end int) vector {
	for i := end; i >= 0; i-- {
		max := getCellMax(v[i])

		for j := i - 1; j >= 0; j-- {
			v[j] = trimCellMax(v[j], max)
		}
	}

	return v
}

func getCellMax(c cell) int {
	var result int

	for i, value := range c.values {
		if value == true {
			result = i
		}
	}

	return result
}

func trimCellMax(c cell, max int) cell {
	for i := max; i < 6; i++ {
		c.values[i] = false
	}

	return c
}

func trimCellMin(c cell, min int) cell {
	for i := 0; i <= min; i++ {
		c.values[i] = false
	}

	return c
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

func logMap(eMap eliminationMap, clues []int) {
	fmt.Println("")

	for i := 0; i < 6; i++ {
		fmt.Printf("         %v  ", clues[i])
	}

	fmt.Println("")

	for i, row := range eMap {
		fmt.Printf("%v   ", clues[23-i])
		logVector(row)
		fmt.Printf("   %v", clues[i+6])

		fmt.Println("")
	}

	for i := 17; i >= 12; i-- {
		fmt.Printf("         %v  ", clues[i])
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
