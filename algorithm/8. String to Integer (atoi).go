package algorithm

import (
	"strings"
)

func myAtoi(str string) int {
	IntMAX := 2147483647
	IntMin := -2147483648

	str = strings.TrimSpace(str)
	if str == "" {
		return 0
	}

	firstEle := string(str[0])
	var firstIndex int
	if firstEle == "-" || firstEle == "+" {
		firstIndex = 1
		if len(str) > 2 {
			if string(str[1]) == "-" || string(str[1]) == "+" {
				return 0
			}
		}
	}

	// if firstEle == "-" {
	// 	if len(str[1:]) > IntMin
	// }

	factor := 1
	l := len(str)
	for i := firstIndex; i < l; i++ {
		if str[i]-'0' > 9 {
			//fmt.Println((str[i] - '0'))
			break
		}
		factor *= 10
	}
	//fmt.Println(factor)
	factor /= 10

	//factor := int(math.Pow(10, float64(l-firstIndex-1)))
	var result int
	for i := firstIndex; i < l; i++ {
		result += int((str[i] - '0')) * factor
		if firstEle == "-" {
			if result > -IntMin/10 || (result == -IntMin && int(str[i]-'0') > -IntMin%10) {
				println(result)
				println(-IntMin)
				return IntMin
			}
		} else {
			if result > IntMAX/10 {
				return IntMAX
			}
		}
		factor /= 10

	}
	if firstEle == "-" {
		result = -result
	}
	return result
}
