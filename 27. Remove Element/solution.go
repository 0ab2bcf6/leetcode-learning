package main

func main() {

}

func removeElement(nums []int, val int) int {

	k, r := 0, len(nums)-1

	for i := 0; i < len(nums); i++ {
		for r > 0 && nums[r] == val {
			r--
		}
		if nums[i] != val {
			k++
		}
		if nums[i] == val && r > i && len(nums) > 1 {
			nums[i] = nums[r]
			nums[r] = val
			k++
		}
	}
	return k
}
