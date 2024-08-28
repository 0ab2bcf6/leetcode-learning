package main

func main() {
}

func maxArea(height []int) int {

	l := 0
	r := len(height) - 1
	var maxArea int = 0

	for l < r {

		if height[l] > height[r] {
			currArea := (r - l) * height[r]
			if currArea > maxArea {
				maxArea = currArea
			}
			r--
		} else {
			currArea := (r - l) * height[l]
			if currArea > maxArea {
				maxArea = currArea
			}
			l++
		}
	}

	return maxArea
}
