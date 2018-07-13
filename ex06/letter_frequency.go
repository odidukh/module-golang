package letter

func Frequency(str string) map[rune]int {
	slice := make([]rune, str)
	var m map[rune]int

	for _, l := range slice {
		if val, ok := m[l]; ok {
			val++
		} else {
			m[l]++
		}
	}
	return m
}

func ConcurrentFrequency(array []string) {
	c := make(chan result)
	c <- Frequency(array[0])
	c <- Frequency(array[1])
	c <- Frequency(array[2])

	select {
	case result := <-c:
		results = append(results, result)
	case <- .....:
		return
	}

}