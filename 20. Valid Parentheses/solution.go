package main

func main() {
}

func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	var stack []rune
	rmap := map[rune]rune{41: 40, 93: 91, 125: 123}

	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
		} else if r == ')' || r == ']' || r == '}' {
			n := len(stack) - 1
			if n < 0 {
				return false
			}

			v, exists := rmap[r]
			if !exists || v != stack[n] {
				return false
			}
			stack = stack[:n]
		}
	}
	return len(stack) == 0
}
