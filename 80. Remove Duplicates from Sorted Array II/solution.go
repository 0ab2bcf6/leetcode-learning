package main

func main() {

}

func removeDuplicates(nums []int) int {

	inc := 0

	for _, num := range nums {
		if inc < 2 || num != nums[inc-2] {
			nums[inc] = num
			inc++
		}
	}

	return inc
}
