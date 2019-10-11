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

	fmt.Println("a", a)
	fmt.Println("maxDenominator", maxDenominator)

	result := maxDenominator

	for i := 1; i < maxDenominator; i++ {
		remainders := 0

		for _, pair := range a {
			remainders += i%pair[0] + i%pair[1]

			if remainders > 0 {
				break
			}
		}

		if remainders == 0 {
			result = i
			break
		}
	}

	fmt.Println("result", result)

	return result
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
		var differential int = denominator / pair[1]
		rebasedNumerator := pair[0] * differential
		rebasedDenominator := pair[1] * differential

		result = append(result, []int{rebasedNumerator, rebasedDenominator})
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
