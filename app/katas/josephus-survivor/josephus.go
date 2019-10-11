package josephus

func JosephusSurvivor(n, k int) int {
	var list []int

	for i := 1; i <= n; i++ {
		list = append(list, i)
	}

	length := len(list)
	index := 0

	for length > 1 {
		index = (index + k - 1) % length

		list = append(list[:index], list[index+1:]...)

		length--
	}

	return list[0]
}
