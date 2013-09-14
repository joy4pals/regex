package main

import "fmt"

func search(pattern, text string) bool {
	if len(pattern) == 0 {
		return true
	} else if pattern[0] == '^' {
		return matchAtBegining(pattern[1:], text)
	} else {
		return matchAtBegining(".*"+pattern, text)
	}
}
func matchAtBegining(pattern, text string) bool {
	if len(pattern) == 0 {
		return true
	} else if pattern[0] == '$' {
		return len(text) == 0
	} else if pattern[1] == '*' {
		return matchStar(pattern[0], pattern[2:], text)
	} else if pattern[0] == '.' || pattern[0] == text[0] {
		return matchAtBegining(pattern[1:], text[1:])
	}
	return false
}
func matchStar(starChar uint8, pattern, text string) bool {
	if matchAtBegining(pattern, text) {
		return true
	} else if starChar == text[0] || starChar == '.' {
		return matchStar(starChar, pattern, text[1:])
	}
	return false
}
func main() {
	fmt.Println(search("ab*", "ababbd"))
	fmt.Println(search("b*", "ababbd"))
	fmt.Println(search("^b*", "ababbd"))
	fmt.Println(search("", "ababbd"))
}
