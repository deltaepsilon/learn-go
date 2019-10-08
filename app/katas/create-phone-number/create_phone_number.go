package create_phone_number

import (
	"fmt"
	"strconv"
)

func CreatePhoneNumber(numbers [10]uint) string {
	var stringNumbers []interface{}

	for _, num := range numbers {
		stringNumbers = append(stringNumbers, strconv.Itoa(int(num)))
	}

	return fmt.Sprintf("(%v%v%v) %v%v%v-%v%v%v%v", stringNumbers...)
}
