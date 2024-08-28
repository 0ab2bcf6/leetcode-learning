
import "fmt"

func main() {
	fmt.Println(intToRoman(58))
}

func intToRoman(num int) string {

	romanValue := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanSymbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""

	for i := 0; i < len(romanValue); i++ {
		for num >= romanValue[i] {
			num -= romanValue[i]
			result += romanSymbol[i]
		}
	}

	return result
}
