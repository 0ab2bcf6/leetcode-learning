package main

func main() {

}

// TODO optimize
func rotate(nums []int, k int) {
	for range k {
		last := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = last
	}
}
