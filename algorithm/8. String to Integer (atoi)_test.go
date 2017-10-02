package algorithm

import (
	"testing"
)

func Test_MyAtoi(t *testing.T) {
	type testCase struct {
		input string
		want  int
	}

	tcs := []testCase{
		testCase{"123", 123},
		testCase{"-123", -123},
		testCase{"1", 1},
		testCase{"-1", -1},
		testCase{"0", 0},
		testCase{"", 0},
		testCase{"12345678910111213", 2147483647},
		testCase{"-12345678910111213", -2147483648},
		testCase{"-", 0},
		testCase{"+", 0},
		testCase{"+-1", 0},
		testCase{"123abv", 123},
		testCase{"-2147483647", -2147483647},
	}

	for _, tc := range tcs {
		out := myAtoi(tc.input)
		if out != tc.want {
			t.Fatalf("input:%s want:%d != out:%d\n", tc.input, tc.want, out)
		}
	}
}
