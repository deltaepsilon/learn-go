package huge

import (
	"fmt"
	"math"
)

func LastDigit(a []int) int {
	var result int
	length := len(a)

	fmt.Println("")
	fmt.Println("a", a)

	if length == 0 {
		result = 1
	} else if length == 1 {
		result = a[0] % 10
	} else {
		result = reduce(a) % 10
	}

	return result
}

func reduce(list []int) int {
	var result int
	lastIndex := len(list) - 1
	a := list[lastIndex-1]
	b := list[lastIndex]
	base := a % 10

	if base == 0 || base == 1 {
		base = a % 100
	}

	result = int(math.Pow(float64(base), float64(b%4)))

	if len(list) > 2 {
		remainder := list[:lastIndex-2]

		result = reduce(append(remainder, result))
	}

	return result
}

// func LastDigit(a []int) int {
// 	var result int
// 	length := len(a)

// 	fmt.Println("")
// 	fmt.Println("a", a)

// 	if length == 0 {
// 		result = 1
// 	} else if length == 1 {
// 		result = a[0] % 10
// 	} else if length == 2 {
// 		base := getBase(a[0])
// 		product := getPower(base, a[1]%4)

// 		result = product % 10

// 		fmt.Println("product, result", product, result)

// 	} else {
// 		var product int = 1

// 		for i := length - 1; i >= 0; i-- {
// 			product = getPower(a[i], product)
// 		}

// 		fmt.Println("product", product)

// 		result = product % 10
// 	}

// 	return result
// }

// func getPower(a, b int) int {
// 	var result int

// 	if a == 0 && b == 0 {
// 		result = 1
// 	} else {
// 		base := getBase(a)
// 		exponent := getExponent(b)
// 		fmt.Println("a,b", a, b)
// 		fmt.Println("base, exponent", base, exponent)
// 		result = int(math.Pow(float64(base), float64(exponent)))

// 	}

// 	return result
// }

// func getBase(n int) int {
// 	base := n % 10

// 	if base == 1 || base == 0 {
// 		base = n % 100
// 	}

// 	return base
// }

// func getExponent(n int) int {
// 	var result int

// 	if n == 0 {
// 		result = 0
// 	} else if n%4 == 0 {
// 		result = 4
// 	} else if n%3 == 0 {
// 		result = 3
// 	} else if n%2 == 0 {
// 		result = 2
// 	} else {
// 		result = 1
// 	}

// 	return result
// }

// func LastDigit(as []int) int {
// 	if len(as) == 0 {
// 		return 1
// 	}

// 	var result int
// 	isSingleDigit := len(as) == 1
// 	lastDigit := getLastDigit(as[0])
// 	isZeroOrFive := lastDigit == 0 || lastDigit == 5

// 	if isSingleDigit {
// 		result = lastDigit
// 	} else if len(as) == 2 {
// 		exponent := getExponent(as[1])
// 		base := getBase(as[0])
// 		product := math.Pow(float64(base), float64(exponent))

// 		result = getLastDigit(int(product))

// 	} else if isZeroOrFive {
// 		result = lastDigit

// 	} else {
// 		var exponent int = 1

// 		for i := len(as) - 1; i >= 0; i-- {
// 			n := as[i]

// 			if n == 0 && exponent == 0 {
// 				exponent = 1
// 				i--
// 			} else {
// 				intermediateExponent := getExponent(exponent)
// 				base := getBase(n)

// 				exponent = int(math.Pow(float64(base), float64(intermediateExponent)))

// 			}
// 		}

// 		result = getLastDigit(exponent)
// 	}

// 	return result
// }

// func getExponent(n int) int {
// 	var result int

// 	if n == 0 {
// 		result = 0
// 	} else if n%4 == 0 {
// 		result = 4
// 	} else if n%2 == 0 {
// 		result = 2
// 	} else if n%3 == 0 {
// 		result = 3
// 	} else {
// 		result = 1
// 	}

// 	return result
// }

// func getBase(n int) int {
// 	base := getLastDigit(n)

// 	if base == 1 || base == 0 && n > 10 {
// 		base = n % 100
// 	}

// 	return base
// }

// func getLastDigit(n int) int {
// 	return n % 10
// }
