package homework

func average(input []float32) (result float32) {
	sum := float32(0)
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	result = sum / float32(len(input))
	return result
}
