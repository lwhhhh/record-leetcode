package algorithm

import "fmt"

func isPalindrome(s string) bool {
	r := []rune(s)
	var i int
	j := len(r) - 1
	for i < j {
		if r[i] != r[j] {
			return false
		}
		i++
		j--
	}
	return true
}

var result [][]string
var oneEles []string

// func backTrace(s string, index int, result *[][]string, oneEles *[]string) {
// 	fmt.Println("len(result1):", len(*result))
// 	if len(s) == index {
// 		*result = append(*result, *oneEles)
// 		//fmt.Println("result:", result, "oneELes:", oneEles)
// 		//fmt.Println("len(result2):", len(result))
// 		return
// 	}
// 	//if len(oneEles) > 0 &&
// 	for i := index; i < len(s); i++ {
// 		t := s[index : i+1]
// 		if !isPalindrome(t) {
// 			continue
// 		}
// 		//fmt.Println(oneEles)
// 		*oneEles = append(*oneEles, t)
// 		backTrace(s, i+1, result, oneEles)
// 		// (*oneEles)[:len(*oneEles)]
// 		for k, e := range *oneEles {
// 			*oneEles[k] =
// 		}
// 	}
// }

func partition(s string) {

}

func integerSplit(n int) int {
	if n == 1 {
		return 1
	}

	for i := 1; i < n; i++ {
		fmt.Printf("%d + ", i)
		//fmt.Printf("%d + ", i)
		//fmt.Printf("i=%d n=%d ", i, n)
		integerSplit(n - i)
		fmt.Printf("%d = %d + %d\n", n, i, n-i)
	}
}
