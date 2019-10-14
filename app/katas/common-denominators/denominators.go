package denominators

import (
	"fmt"
)

func ConvertFracts(a [][]int) string {
	denominator := getCommonDenominator(a)
	rebased := rebaseFractions(a, denominator)

	return stringifyResult(rebased)
}

func getCommonDenominator(a [][]int) int {
	maxDenominator := getMaxDenominator(a)
	maxJ := len(a) - 1

	for i := 1; i < maxDenominator; i++ {
		for j, pair := range a {
			remainder := i * pair[0] % pair[1]
			isLast := j == maxJ

			if remainder > 0 {
				break
			} else if isLast {
				return i
			}
		}
	}

	return maxDenominator
}

func getMaxDenominator(a [][]int) int {
	result := 1

	for _, pair := range a {
		result *= pair[1]
	}

	return result
}

func rebaseFractions(a [][]int, denominator int) [][]int {
	result := [][]int{}

	for _, pair := range a {
		rebasedNumerator := denominator * pair[0] / pair[1]

		result = append(result, []int{rebasedNumerator, denominator})
	}

	return result
}

func stringifyResult(a [][]int) string {
	result := ""

	for _, pair := range a {
		result += fmt.Sprintf("(%v,%v)", pair[0], pair[1])
	}

	return result
}
