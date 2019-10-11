package squares

import (
	"math"
)

func SquaresInRect(length int, width int) []int {
	if length == width {
		return nil
	}

	shortSide := int(math.Min(float64(length), float64(width)))
	remainder := int(math.Abs(float64(length) - float64(width)))
	result := []int{shortSide}

	if remainder > 0 {
		nextResult := SquaresInRect(shortSide, remainder)

		result = append(result, nextResult...)

	}

	if remainder == shortSide {
		result = append(result, shortSide)
	}

	return result
}
