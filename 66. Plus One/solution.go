package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9, 9}))
}

func plusOne(digits []int) []int {

	i := len(digits) - 1
	carry := 1

	for carry != 0 {
		if i == 0 && digits[i] == 9 {
			digits[i] = 0
			digits = append(digits, 0)
			for j := len(digits) - 1; j >= 0; j-- {
				if j != 0 {
					digits[j] = digits[j-1]
				} else {
					digits[j] = 1
				}
			}
		} else {
			if digits[i] == 9 {
				digits[i] = 0
				i--
				carry++
			} else {
				digits[i] = digits[i] + 1
			}
		}
		carry--
	}
	return digits
}
