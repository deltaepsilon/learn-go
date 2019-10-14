package twice

// Solution is waaaaay too slow
func DblLinear(n int) int {
	i := 0
	x := 1
	u := []int{x}

	for i < n {
		y := 2*x + 1
		z := 3*x + 1
		u = appendInOrder(u, y)
		u = appendInOrder(u, z)

		i++
		x = u[i]
	}

	return u[n]
}

func appendInOrder(u []int, x int) []int {
	a := make([]int, len(u))
	copy(a, u)

	i := len(a)

	for i > 0 {
		i--

		if a[i] == x {
			break
		} else if a[i] < x {
			a = append(a[:i+1], append([]int{x}, a[i+1:]...)...)
			break
		}
	}

	return a
}
