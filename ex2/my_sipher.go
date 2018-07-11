package cipher

type Caesar struct {
}

type Shift struct {
	num int
}

type Vigenere struct {
	str string
}

func NewCaesar() Cipher {
	var caesar Cipher = &Caesar{}
	return caesar

}

func (cipher Caesar) Encode(str string) string {
	c := NewShift(3)
	ciphered := c.Encode(str)
	return ciphered
}

func (cipher Caesar) Decode(str string) string {
	c := NewShift(3)
	deciphered := c.Decode(str)
	return deciphered
}

func NewShift(distance int) Cipher {
	var shift Cipher
	switch {
	case distance > 0 && distance < 26:
		shift = &Shift{distance}
		return shift
	case distance > -26 && distance < 0:
		distance += 26
		shift = &Shift{distance}
		return shift
	}
	return nil
}

func (cipher Shift) Encode(str string) string {
	var plain []rune = []rune(CorrectString(str))
	ciphered := make([]rune, 0)

	for i := 0; i < len(plain); i++ {
		res := rune((int(plain[i])-97+cipher.num)%26 + 97)
		if res > 96 {
			ciphered = append(ciphered, res)
		} else {
			res = rune((int(plain[i])-97+cipher.num)%26 + 123)
			ciphered = append(ciphered, res)
		}
	}
	return string(ciphered)
}

func (cipher Shift) Decode(str string) string {
	var ciphered []rune = []rune(str)
	deciphered := make([]rune, 0)

	for i := 0; i < len(ciphered); i++ {
		r := ciphered[i] - rune(cipher.num)
		if r > 96 && r < 123 {
			deciphered = append(deciphered, r)
		} else if r > 122 {
			r = ciphered[i] - rune(26+cipher.num)
			deciphered = append(deciphered, r)
		} else if r < 96 {
			r = ciphered[i] + rune(26-cipher.num)
			deciphered = append(deciphered, r)
		}
	}
	return string(deciphered)
}

func NewVigenere(key string) Cipher {
	if NonZero(key) == 0 || key != CorrectString(key) || len(key) == 0 {
		return nil
	} else {
		var vigenere Cipher = &Vigenere{key}
		return vigenere
	}
}

func (cipher Vigenere) Encode(str string) string {
	var text []rune = []rune(CorrectString(str))
	ciphered := make([]rune, 0)

	for i := 0; i < len(text); i++ {
		shift := NewShift(int(cipher.str[i%len(cipher.str)] - 97))
		var nextstr string
		if shift != nil {
			nextstr = shift.Encode(string(text[i]))
		} else {
			nextstr = string(text[i])
		}
		next := []rune(nextstr)
		ciphered = append(ciphered, next[0])
	}
	return string(ciphered)
}

func (cipher Vigenere) Decode(str string) string {
	var ciphered []rune = []rune(str)
	decoded := make([]rune, 0)

	for i := 0; i < len(str); i++ {
		var nextstr string
		shift := NewShift(-int(cipher.str[i%len(cipher.str)] - 97))
		if shift != nil {
			nextstr = shift.Encode(string(ciphered[i]))
		} else {
			nextstr = string(ciphered[i])
		}
		next := []rune(nextstr)
		decoded = append(decoded, next[0])
	}
	return string(decoded)
}

func CorrectString(str string) string {
	var str_slice []rune = []rune(str)
	var char_array []rune = make([]rune, 0)

	for i := 0; i < len(str); i++ {
		if str_slice[i] > 64 && str_slice[i] < 91 {
			str_slice[i] += 32
			char_array = append(char_array, str_slice[i])

		} else if str_slice[i] > 96 && str_slice[i] < 123 {
			char_array = append(char_array, str_slice[i])
		}
	}
	return string(char_array)
}

func NonZero(str string) int {
	var sum int
	for i := 0; i < len(str); i++ {
		sum += int(str[i])
	}

	if len(str) == 0 || sum/len(str) == 97 {
		return 0
	} else {
		return 1
	}

}
