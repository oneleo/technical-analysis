package calculate

func Average(data []float64) (result float64) {

	result = 0.0
	length := len(data)

	for i := 0; i < length; i++ {
		result = result + data[i]
	}

	result = result / float64(length)

	return result
}
