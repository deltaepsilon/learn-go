package partitions

import (
	"math"
)

func Partitions(n int) int {
	var lag int = 0

	if n > 1 {
		lag = Partitions(n - 1)
	}

	partitions := getPartitions(n, [][]int{{n}})

	return lag + len(partitions)
}

func getPartitions(n int, partitions [][]int) [][]int {
	result := partitions

	for _, partition := range partitions {
		for i, digit := range partition {
			tail := partition[i+1:]
			parts := getParts(digit)

			if len(parts) == 0 {
				break
			}

			var subPartitions [][]int

			for i := 0; i < len(parts); i++ {
				part := parts[i]
				partMin := mapMath(getFloat(part), math.Min)
				tailMax := mapMath(getFloat(tail), math.Max)
				isMatchingSum := n == sumSlice(part)+sumSlice(tail)
				isOrdered := partMin >= tailMax

				if isMatchingSum && isOrdered {
					subPartitions = append(subPartitions, append(part, tail...))
				}
			}

			result = append(result, getPartitions(n, subPartitions)...)

		}
	}

	return result
}

func getParts(n int) [][]int {
	var parts [][]int

	for i := 1; i < n; i++ {
		part := n - i

		if part >= i && i != 1 {
			parts = append(parts, []int{part, i})
		}
	}

	return parts
}

func sumSlice(s []int) int {
	var result int = 0

	for _, n := range s {
		result += n
	}

	return result
}

func getFloat(integers []int) []float64 {
	result := make([]float64, len(integers))

	for i, n := range integers {
		result[i] = float64(n)
	}

	return result
}

func mapMath(numbers []float64, f func(float64, float64) float64) float64 {
	var result float64

	if len(numbers) > 0 {
		result = numbers[0]
	}

	for i := 1; i < len(numbers); i++ {
		result = f(result, numbers[i])
	}

	return result
}
