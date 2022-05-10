package homework

func average(input [15]float32) (result float32) {
	var sum float32
	arrayLength := len(input)
	for i := 0; i < arrayLength; i++ {
		sum += input[i]
	}
	result = sum / float32(arrayLength)
	return
}
