package letter

func Frequency(text string) map[rune]int {
	total := make(map[rune]int)

	for _, count := range text {
		total[count]++
	}

	return total
}

func ConcurrentFrequency(text []string) map[rune]int {
	channel := make(chan map[rune]int)

	for _, list := range text {
		go func(word string) {
			channel <- Frequency(word)
		}(list)
	}

	total := make(map[rune]int)

	for range text {
		for key, value := range <-channel {
			total[key] += value
		}
	}
	return total
}
