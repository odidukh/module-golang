package brackets

func Bracket(str string) (bool, error) {
	stack := New()
	var slice []rune = []rune(str)

	for _, t := range slice {
		if t == '{' || t == '[' || t == '(' {
			stack.Push(t)
		} else {
			s := stack.Pop()
			switch {
			case (t == '}' && s == '{'):
				continue
			case (t == ']' && s == '['):
				continue
			case (t == ')' && s == '('):
				continue
			default:
				return false, nil
			}
		}
	}
	if stack.Len() != 0 {
		return false, nil
	}
	return true, nil
}
