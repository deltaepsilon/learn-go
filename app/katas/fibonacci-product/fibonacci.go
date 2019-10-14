package fibonacci

func ProductFib(prod uint64) [3]uint64 {
	var n0 uint64 = 0
	var n1 uint64 = 1
	var found uint64 = 0
	var product uint64 = 0

	for product <= prod {
		n1 = n0 + n1
		n0 = n1 - n0
		product = n0 * n1

		if product == prod {
			found = 1
			break
		}
	}

	return [3]uint64{n0, n1, found}
}
