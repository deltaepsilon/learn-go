package reverse_rotate

import (
	"math"
	"strconv"
)

func Revrot(s string, n int) string {
	if n == 0 {
		return ""
	}

	if s == "" {
		return ""
	}

	var result string

	for _, chunk := range getChunks(s, n) {
		sumOfCubes := calcSumOfCubes(chunk)
		isDivisibleByTwo := getIsDivisibleByTwo(sumOfCubes)

		if isDivisibleByTwo {
			result += reverse(chunk)
		} else {
			result += rotate(chunk)
		}
	}

	return result
}

func getChunks(s string, n int) []string {
	var wholeChunkCount int = len(s) / n
	var chunks []string

	for i := 0; i < wholeChunkCount; i++ {
		startIndex := i * n
		endIndex := startIndex + n

		chunks = append(chunks, s[startIndex:endIndex])
	}

	return chunks
}

func calcSumOfCubes(s string) int {
	var sumOfCubes int = 0

	for _, r := range s {
		digit, _ := strconv.Atoi(string(r))
		cube := int(math.Pow(float64(digit), float64(3)))

		sumOfCubes += cube
	}

	return sumOfCubes
}

func getIsDivisibleByTwo(i int) bool {
	return i%2 == 0
}

func reverse(s string) string {
	var result string = ""

	for _, r := range s {
		result = string(r) + result
	}

	return result
}

func rotate(s string) string {
	return s[1:] + s[0:1]
}
