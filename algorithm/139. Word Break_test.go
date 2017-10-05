package algorithm

import (
	"testing"
)

func Test_wordBreak(t *testing.T) {
	s := "abcd"
	ss := []string{"a", "abc", "b", "cd"}
	wordBreak(s, ss)

	t.Logf("\n")

	s = "a"
	ss = []string{"a"}
	wordBreak(s, ss)

	t.Logf("\n")

	s = "leetcode"
	ss = []string{"leet", "code"}
	wordBreak(s, ss)
}
