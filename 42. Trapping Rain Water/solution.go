package main

func main() {
}

func trap(height []int) int {

	// get the max value in the array
	// while max > 1
	// check if there are more than 2 of the maxvalue
	// if yes > count horizontal the gaps between max and max-1
	// max--
	max := 0
	for _, i := range height {
		if i > max {
			max = i
		}
	}

	result := 0

	for max > 0 {
		for i, curr := range height {

		}
		max--
	}

	return 0
}