package digital_root

import (
	"strconv"
)

func DigitalRoot(n int) int {
	result := n

	if n >= 10 {
		var sum int = 0

		for _, digit := range strconv.Itoa(n) {
			i, _ := strconv.Atoi(string(digit))

			sum += i
		}

		result = DigitalRoot(sum)
	}

	return result
}
