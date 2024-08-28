package main

import "sort"

func main() {

}

func hIndex(citations []int) int {

	// can be improved with inbuilt sorting method
	sort.Slice(citations, func(i, j int) bool { return citations[i] > citations[j] })
	count := 0

	for _, c := range citations {
		if count < c {
			count++
		} else {
			break
		}
	}
	return count
}
