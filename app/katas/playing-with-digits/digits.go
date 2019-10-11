package digits

import (
	"math"
	"strconv"
)

func DigPow(n, p int) int {
	digits := getDigits(n)
	sum := getExponentialSum(digits, p)
	result := sum / n
	remainder := sum % n

	if remainder != 0 {
		result = -1
	}

	return result
}

func getDigits(n int) []int {
	var result []int

	for _, digit := range strconv.Itoa(n) {
		i, _ := strconv.Atoi(string(digit))
		result = append(result, i)
	}

	return result
}

func getExponentialSum(digits []int, p int) int {
	result := 0

	for i, digit := range digits {
		result += int(math.Pow(float64(digit), float64(i+p)))
	}

	return result
}
