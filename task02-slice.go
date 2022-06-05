package homework

func reverse(input []int64) (result []int64) {
	// fmt.Printf("(%T:%+v)", input, input)
	for i := len(input) - 1; i >= 0; i-- {
		result = append(result, input[i])
	}
	// fmt.Printf("%+v", result)
	return result
}
