package huge

import (
	"math"
)

func LastDigit(as []int) int {
	if len(as) == 0 {
		return 1
	}

	var result int
	isSingleDigit := len(as) == 1
	lastDigit := getLastDigit(as[0])
	isZeroOrFive := lastDigit == 0 || lastDigit == 5

	if isSingleDigit {
		result = lastDigit
	} else if len(as) == 2 {
		exponent := getExponent(as[1])
		base := getBase(as[0])
		product := math.Pow(float64(base), float64(exponent))

		result = getLastDigit(int(product))

	} else if isZeroOrFive {
		result = lastDigit

	} else {
		var exponent int = 1

		for i := len(as) - 1; i >= 0; i-- {
			n := as[i]

			if n == 0 && exponent == 0 {
				exponent = 1
				i--
			} else {
				intermediateExponent := getExponent(exponent)
				base := getBase(n)

				exponent = int(math.Pow(float64(base), float64(intermediateExponent)))

			}
		}

		result = getLastDigit(exponent)
	}

	return result
}

func getExponent(n int) int {
	var result int

	if n == 0 {
		result = 0
	} else if n%4 == 0 {
		result = 4
	} else if n%2 == 0 {
		result = 2
	} else if n%3 == 0 {
		result = 3
	} else {
		result = 1
	}

	return result
}

func getBase(n int) int {
	base := getLastDigit(n)

	if base == 1 || base == 0 && n > 10 {
		base = n % 100
	}

	return base
}

func getLastDigit(n int) int {
	return n % 10
}
