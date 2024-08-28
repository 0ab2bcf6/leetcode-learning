package main

func main() {
	_ = letterCombinations("23")
}

func letterCombinations(digits string) []string {

	var result []string
	letters := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	for i, input_r := range digits {
		index := int(input_r-'0') - 2
		if i == 0 {
			for _, char := range letters[index] {
				new_char = string(char)
				new_slice := []string(new_char)
				result = append(result)
			}
		} else {

		}
	}
	return result
}
