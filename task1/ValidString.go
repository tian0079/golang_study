package task1

func ValidString(str string) bool {

	slice := []string{}

	maps := map[string]string{"(": ")", "[": "]", "{": "}"}

	for _, chars := range str {
		char, ok := maps[string(chars)]
		if ok {
			slice = append(slice, char)
		} else {
			if len(slice) == 0 {
				return false
			}
			v := slice[len(slice)-1]
			slice = slice[:len(slice)-1]
			if v != string(chars) {
				return false
			}
		}

	}
	return len(slice) == 0
}
