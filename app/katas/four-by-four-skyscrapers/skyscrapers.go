package skyscrapers

import (
	"fmt"
	"strconv"
)

type eliminationMap map[int]bool
type eliminationGrid [4][4]eliminationMap

func SolvePuzzle(clues []int) [][]int {
	if len(clues) == 0 {
		return [][]int{{0}}
	}

	eGrid := getEliminationGrid()
	eGrid = processFours(clues, eGrid)
	eGrid = processOnes(clues, eGrid)
	eGrid = processTwos(clues, eGrid)
	eGrid = processThrees(clues, eGrid)

	return getResultFromGrid(eGrid)
}

func getEliminationGrid() eliminationGrid {
	var result eliminationGrid

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = map[int]bool{1: true, 2: true, 3: true, 4: true}
		}
	}

	return result
}

func processFours(clues []int, eGrid eliminationGrid) eliminationGrid {
	for i, clue := range clues {
		if clue == 4 {
			vector := getClueVector(i, eGrid)

			for mapIndex, eMap := range vector {
				for key := range eMap {
					vector[mapIndex][key] = key == mapIndex+1
				}
			}
		}
	}

	return eliminate(eGrid)
}

func processOnes(clues []int, eGrid eliminationGrid) eliminationGrid {
	for i, clue := range clues {
		if clue == 1 {
			vector := getClueVector(i, eGrid)

			vector[0][1] = false
			vector[0][2] = false
			vector[0][3] = false
		}

	}

	return eliminate(eGrid)
}

func processTwos(clues []int, eGrid eliminationGrid) eliminationGrid {
	for i, clue := range clues {
		if clue == 2 {
			vector := getClueVector(i, eGrid)

			// The first item can't be a 4
			vector[0][4] = false

			// Items between first and next possible 4 must be lower than first item
			maxHeightOfFirst := getMaxHeight(vector[0])

			for i := 1; i < 4; i++ {
				if vector[i][4] == true {
					break
				} else {
					eliminateGreaterThan(vector[i], maxHeightOfFirst-1)
				}
			}
		}
	}

	getHasChanged := getGetHasChanged(eGrid)
	result := eliminate(eGrid)
	hasChanged := getHasChanged(result)

	if hasChanged {
		result = processTwos(clues, result)
	}

	return result
}

func processThrees(clues []int, eGrid eliminationGrid) eliminationGrid {
	for i, clue := range clues {

		if clue == 3 {
			vector := getClueVector(i, eGrid)

			// The first two items can't be 4
			vector[0][4] = false
			vector[1][4] = false

			// Items between first and next possible 4 must be lower than first item
			maxHeightOfFirst := getMaxHeight(vector[0])
			maxHeightOfSecond := getMaxHeight(vector[1])
			maxHeightOfThird := getMaxHeight(vector[2])
			maxHeightOfFourth := getMaxHeight(vector[3])

			if maxHeightOfThird == 4 {
				// trim the first item if there's no room to rise to the highest item
				if maxHeightOfFirst >= maxHeightOfSecond {
					vector[0][maxHeightOfFirst] = false
				}
			} else if maxHeightOfFourth == 4 {
				// trim the first item if there's no room to rise to the highest item
				if maxHeightOfFirst >= maxHeightOfSecond && maxHeightOfFirst >= maxHeightOfThird {
					vector[0][maxHeightOfFirst] = false
				}
			}
		}
	}

	getHasChanged := getGetHasChanged(eGrid)
	result := eliminate(eGrid)
	hasChanged := getHasChanged(result)

	if hasChanged {
		result = processThrees(clues, result)
	}

	return result
}

func eliminate(eGrid eliminationGrid) eliminationGrid {
	getHasChanged := getGetHasChanged(eGrid)
	result := executeElimination(eGrid)
	hasChanged := getHasChanged(result)

	if hasChanged {
		result = eliminate(result)
	}

	return result
}

func executeElimination(eGrid eliminationGrid) eliminationGrid {
	columns := getColumns(eGrid)

	for _, row := range eGrid {
		eliminateTrueSiblings(row)
		promoteFound(row)
	}

	for _, column := range columns {
		eliminateTrueSiblings(column)
		promoteFound(column)
	}

	return eGrid
}

func eliminateTrueSiblings(vector [4]eliminationMap) {
	for i := 0; i < 4; i++ {
		eMap := vector[i]
		trueKeys := getTrueKeys(eMap)

		if len(trueKeys) == 1 {
			eliminateKey(vector, trueKeys[0], i)
		}
	}
}

func promoteFound(vector [4]eliminationMap) {
	for height := 1; height <= 4; height++ {
		var count int
		var lastTrueIndex int

		for i := 0; i < 4; i++ {
			if vector[i][height] == true {
				count++
				lastTrueIndex = i
			}
		}

		if count == 1 {
			for key := range vector[lastTrueIndex] {
				vector[lastTrueIndex][key] = key == height
			}
		}
	}
}

func getClueVector(i int, eGrid eliminationGrid) [4]eliminationMap {
	var clueVector [4]eliminationMap

	if i < 4 {

		columns := getColumns(eGrid)
		column := columns[i%4]

		for i := 0; i < 4; i++ {
			clueVector[3-i] = column[i]
		}
	} else if i < 8 {
		row := eGrid[i%4]

		for i := 0; i < 4; i++ {
			clueVector[3-i] = row[i]
		}
	} else if i < 12 {
		columns := getColumns(eGrid)
		column := columns[3-i%4]

		clueVector = column
	} else {
		clueVector = eGrid[3-i%4]
	}

	return clueVector
}

func getTrueKeys(eMap eliminationMap) []int {
	var keys []int

	for key, value := range eMap {
		if value == true {
			keys = append(keys, key)
		}
	}

	return keys
}

func getColumns(eGrid eliminationGrid) eliminationGrid {
	var columnarGrid [4][4]eliminationMap

	for i := 0; i < 4; i++ {

		for j := 0; j < 4; j++ {
			columnarGrid[i][3-j] = eGrid[j][i]
		}
	}

	return columnarGrid
}

func eliminateKey(row [4]eliminationMap, key int, indexToSkip int) [4]eliminationMap {
	for i, eMap := range row {
		if i != indexToSkip {
			eMap[key] = false
		}
	}

	return row
}

func getMaxHeight(eMap eliminationMap) int {
	var result int

	for i := 4; i > 0; i-- {
		if eMap[i] == true {
			result = i
			break
		}
	}

	return result
}

func eliminateGreaterThan(eMap eliminationMap, max int) {
	for i := 4; i > max; i-- {
		eMap[i] = false
	}
}

func getGetHasChanged(beforeGrid eliminationGrid) func(eliminationGrid) bool {
	beforeSignature := fmt.Sprint(beforeGrid)

	return func(afterGrid eliminationGrid) bool {
		afterSignature := fmt.Sprint(afterGrid)

		return beforeSignature != afterSignature
	}
}

func getResultFromGrid(eGrid eliminationGrid) [][]int {
	var result [][]int

	for rowIndex, row := range eGrid {
		result = append(result, []int{0, 0, 0, 0})

		for columnIndex, eMap := range row {
			for key, value := range eMap {
				if value == true {
					result[rowIndex][columnIndex] = key
					break
				}
			}
		}
	}

	return result
}

func printGrid(eGrid eliminationGrid) {
	for _, row := range eGrid {
		printVector(row)
	}

	fmt.Println("")
	fmt.Println("")
}

func printVector(vector [4]eliminationMap) {
	var pretty [4][4]string

	for i, eliminationMap := range vector {
		for key, value := range eliminationMap {
			if value == true {
				pretty[i][key-1] = strconv.Itoa(key)
			} else if key > 0 {
				pretty[i][key-1] = " "
			} else {
				fmt.Println("out key:", key)

			}
		}
	}
	fmt.Println(pretty)
}
