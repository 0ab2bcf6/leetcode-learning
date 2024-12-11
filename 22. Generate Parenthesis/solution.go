package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {

	result := make([]string, 0)

	genRecursive(&result, 2*n, 0, "", 0)

	return result
}

func genRecursive(lst *[]string, depth int, i int, str string, openbrckts int) {
	if depth == 2 {
		*lst = append(*lst, "()")
		return
	}

	if depth == i {
		*lst = append(*lst, str)
		return
	}

	if openbrckts < (depth/2) && (depth-i) >= 2 && openbrckts < depth-i {
		genRecursive(lst, depth, i+1, str+"(", openbrckts+1)
	}
	if openbrckts > 0 {
		genRecursive(lst, depth, i+1, str+")", openbrckts-1)
	}
}
