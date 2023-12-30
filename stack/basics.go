package stack

func AreBracketsValid(brackets string) bool {
	bracketsArr := []byte(brackets)
	stack := Stack{}
	for _, b := range bracketsArr {
		bracket := string(b)
		if isOpeningBracket(bracket) {
			stack.Push(bracket)
		} else {
			bracketFromStack := (stack.Pop()).(string)
			if bracketFromStack == "" || !areSameBracket(bracketFromStack, bracket) {
				return false
			}
		}
	}
	return true
}

func isOpeningBracket(bracket string) bool {
	return bracket == "{" || bracket == "(" || bracket == "["
}

func areSameBracket(opening string, closing string) bool {
	return (opening == "{" && closing == "}") || (opening == "(" && closing == ")") || opening == "[" && closing == "]"
}
