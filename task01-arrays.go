package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var total float32 = 0
	for _, number := range input {
		total += number
	}
	return total / float32(len(input))
}
