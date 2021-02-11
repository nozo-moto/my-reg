package myreg

// func main() {
// 	fmt.Println(match("a", "a"))
// 	fmt.Println(match("a", "b"))
// 	fmt.Println(match("", ""))
// 	fmt.Println(match(".", "a"))
// 	fmt.Println(match("", "abc"))
// }

func match(pattern, text string) bool {
	if pattern == "" {
		return true
	} else if pattern == "$" && text == "" {
		return true
	} else if pattern == "?" {
		return matchQuestion(pattern, text)
	} else if pattern == "*" {
		return matchStar(pattern, text)
	}

	return matchOne(string(pattern[0]), string(text[0])) && match(pattern[1:], text[1:])
}

func search(pattern, text string) bool {
	if string(pattern[0]) == "^" {
		return match(pattern[1:], text)
	}

	return match(".*"+pattern, text)
	// for i := range text {
	// 	if match(pattern, text[i:]) {
	// 		return true
	// 	}
	// }
	// return false
}

func matchOne(pattern, text string) bool {
	if pattern == "" {
		return true
	}
	if text == "" {
		return false
	}

	if pattern == "." {
		return true
	}

	return pattern == text
}

func matchQuestion(pattern, text string) bool {
	if matchOne(string(pattern[0]), string(text[0])) && match(pattern[2:], text[1:]) || match(pattern[2:], text) {
		return true
	}
	return match(pattern[2:], text)
}

func matchStar(pattern, text string) bool {
	return (matchOne(string(pattern[0]), string(text[0])) && match(pattern, text[1:])) || match(pattern[2:], text)
}
