package reverse_rotate

import (
	"fmt"
	"strconv"
	"strings"
)

func Revrot(s string, n int) string {
	if n == 0 {
		return ""
	}

	if s == "" {
		return ""
	}

	chunks := getChunks(s, n)

	reversedChunks := chunks

	var stringChunks []string

	for _, chunk := range reversedChunks {
		stringChunks = append(stringChunks, strconv.Itoa(chunk))
	}

	fmt.Println(reversedChunks, stringChunks)

	// for _, digit := range s {

	// }

	return strings.Join(stringChunks, "")
}

func getChunks(s string, n int) []int {
	var wholeChunkCount int = len(s) / n
	var remainderChunkCount int = len(s) % n
	var chunks []int

	for i := 0; i < wholeChunkCount; i++ {
		var num, _ = strconv.Atoi(string(s[i*n : i*n+n]))

		chunks = append(chunks, num)
		fmt.Println(i, chunks, num)
	}

	if remainderChunkCount > 0 {
		var num, _ = strconv.Atoi(string(s[wholeChunkCount*n : len(s)]))

		chunks = append(chunks, num)
	}

	return chunks
}
