package huge

import (
	"fmt"
	"math"
)

func LastDigit(as []int) int {
	if len(as) == 0 {
		return 1
	}

	var result int
	lastDigit := getLastDigit(as[0])

	if len(as) == 1 {
		result = int(as[0]) % 10

	} else if len(as) == 2 {
		power := getBase(as[1])
		product := math.Pow(float64(as[0]), float64(power))
		result = int(product) % 10

	} else if lastDigit == 0 || lastDigit == 5 {
		result = lastDigit
	} else {
		var power float64 = 1

		fmt.Println("as", as)

		for i := len(as) - 1; i > 0; i-- {
			n := float64(as[i])

			if n == 0 && power == 0 {
				power = 1
				i--
			} else {
				raised := math.Pow(n, power)
				base := getBase(int(raised))
				power = float64(base)
			}
		}

		product := as[0]

		
		fmt.Println("as[0]", as[0])
		
		for j := power - 1; j > 0; j-- {
			
			product *= as[0]
			fmt.Println("product", product)
			
		}

		result = getLastDigit(product)

		// fmt.Println("as[0], power", as[0], power)
		// fmt.Println("float64(as[0])", float64(as[0]))

		// product = math.Pow(float64(as[0]), power)

		// fmt.Println("product", product)
		// fmt.Println("int(product)", int(product))

		// result = int(product) % 10

		// fmt.Println("result", result)

	}

	return result
}

func getBase(n int) int {
	var result int

	if n == 0 {
		result = 0
	} else if n%4 == 0 {
		result = 4
	} else if n%3 == 0 {
		result = 3
	} else if n%2 == 0 {
		result = 2
	} else {
		result = 1
	}

	return result
}

func getLastDigit(n int) int {
	return n % 10
}
