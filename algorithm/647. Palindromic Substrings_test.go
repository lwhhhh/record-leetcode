package algorithm

import "testing"

func Test_PalindromicSubstrings(t *testing.T) {
	testCase := make(map[string]int)

	testCase["aaa"] = 6
	testCase["abc"] = 3
	testCase["a"] = 1

	for k, v := range testCase {
		out := countSubstrings(k)
		if out != v {
			t.Fatal("out != want:", out, v)
		}
	}
}
