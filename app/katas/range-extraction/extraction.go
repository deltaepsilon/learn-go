package extraction

import (
	"strconv"
)

func Solution(list []int) string {
	var result string = strconv.Itoa(list[0])

	for i := 0; i < len(list); i++ {
		number := list[i]
		lastConsecutive := getLastConsecutive(number, i, list)
		diff := lastConsecutive - number

		if i > 0 {
			result += "," + strconv.Itoa(number)
		}

		if diff > 1 {
			result += "-" + strconv.Itoa(lastConsecutive)

			i += diff
		}

	}

	return result
}

func getLastConsecutive(number, startIndex int, list []int) int {
	result := number

	for j := startIndex + 1; j < len(list); j++ {
		if list[j]-result == 1 {
			result = list[j]
		} else {
			break
		}
	}

	return result
}
