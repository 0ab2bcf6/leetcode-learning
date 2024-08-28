package main

import "fmt"

func main() {
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reversed := 0
	y := x

	for y > 0 {
		reversed = reversed*10 + y%10
		y /= 10
	}

	return (reversed == x)
}

func isPalindrome_bad(x int) bool {
	runeArray := []rune(fmt.Sprintf("%d", x))
	lenArray := len(runeArray)

	if lenArray%2 == 0 {
		j := lenArray - 1
		for i := 0; i < lenArray/2; i++ {
			if runeArray[i] != runeArray[j] {
				return false
			}
			j--
		}
	} else {
		j := lenArray - 1
		for i := 0; i < (lenArray / 2); i++ {
			if runeArray[i] != runeArray[j] {
				return false
			}
			j--
		}
	}
	return true
}
