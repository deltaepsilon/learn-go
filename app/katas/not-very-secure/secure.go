package secure

import (
	"regexp"
)

func Alphanumeric(s string) bool {
	result, _ := regexp.Match("^[a-zA-Z0-9]+$", []byte(s))

	return result
}
