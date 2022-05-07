package homework

func reverse(input []int64) (result []int64) {
	for n := len(input) - 2; n >= 0; n-- {
		result = append(input[:n], append(input[n+1:], input[n])...)
	}
	return result
}
