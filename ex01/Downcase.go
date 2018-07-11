package downcase

func Downcase(str string) (string, error) {
	var sl []rune
	sl = []rune(str)
	for i := 0; i < len(sl); i++ {
		if sl[i] >= 65 && sl[i] <= 90 {
			sl[i] += 32
		}
	}
	str = string(sl)
	return str, nil
}
