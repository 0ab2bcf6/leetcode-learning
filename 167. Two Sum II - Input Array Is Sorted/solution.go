package main

func main() {

}

func twoSum(numbers []int, target int) []int {

	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if target-sum == 0 {
			return []int{l + 1, r + 1}
		}
		if target-sum > 0 {
			r--
		} else {
			l++
		}
	}
	return []int{-1, -1}
}

// naive
// func twoSum(numbers []int, target int) []int {
// 	for i := 0; i < len(numbers); i++ {
// 		for j := i + 1; j < len(numbers); j++ {
// 			if numbers[i]+numbers[j] == target {
// 				return []int{i+1, j+1}
// 			}
// 		}
// 	}
// 	return []int{-1, -1}
// }
