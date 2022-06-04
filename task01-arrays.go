package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0
	var count int = 0
	for i, el := range input {
		if el != 0 {
			count++
			sum += el
			println(i, el)
		}
	}
	// println(sum, " ", count)
	return sum / float32(count)
}
