package main

func main() {

}

func convert(s string, numRows int) string {

	runeArray := make([][]rune, numRows)
	var result string = ""
	var step int = 1

	index := 0

	for _, c := range s {
		//fmt.Println(i, index)
		runeArray[index] = append(runeArray[index], c)

		if numRows == 1 {
			step = 0
		} else if index+step == numRows {
			step = -1
		} else if index+step == -1 {
			step = 1
		}
		index += step
	}

	for _, runes := range runeArray {
		for _, r := range runes {
			result += string(r)
		}
	}

	return result
}
