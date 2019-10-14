package sudoku

func ValidateSolution(rows [][]int) bool {
	return validateSets(rows) &&
		validateSets(getColumns(rows)) &&
		validateSets(getGrids(rows))
}

func validateSets(sets [][]int) bool {
	for _, set := range sets {
		product := 1

		for _, value := range set {
			product *= value
		}

		if product != 362880 {
			return false
		}
	}

	return true
}

func getColumns(rows [][]int) [][]int {
	columns := [][]int{}

	for _, row := range rows {
		for j, columnValue := range row {

			if len(columns)-1 < j {
				columns = append(columns, []int{})
			}

			columns[j] = append(columns[j], columnValue)
		}
	}

	return columns
}

func getGrids(rows [][]int) [][]int {
	grids := [][]int{}

	for i, row := range rows {
		for j, columnValue := range row {
			iQuotient := i / 3
			jQuotient := j / 3
			row := jQuotient + 3*iQuotient

			if len(grids)-1 < row {
				grids = append(grids, []int{})
			}

			grids[row] = append(grids[row], columnValue)
		}
	}

	return grids
}
