package homework

func average(input []float32) (result float32) {
	var sum float32 = 0
	for i, el := range input {
		sum += el
		println(i, el)
	}
	println(sum, " ", len(input))
	return sum / float32(len(input))
}
