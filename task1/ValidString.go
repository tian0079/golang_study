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

	// for i := 0; i < len(str); i++ {
	// 	s := string(str[i])
	// 	if s == "(" || s == "[" || s == "{" {
	// 		slice = append(slice, maps[s])
	// 	} else {

	// 		if len(slice) == 0 {
	// 			return false
	// 		}
	// 		v := slice[len(slice)-1]
	// 		slice = slice[:len(slice)-1]
	// 		if v != s {
	// 			return false
	// 		}

	// 	}

	// }

	return len(slice) == 0
}
