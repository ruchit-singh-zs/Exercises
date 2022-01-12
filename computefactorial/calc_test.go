package factorial

import (
	"testing"
)

/*
	Test Driven Table
*/

func TestFactorial(t *testing.T) {
	cases := []struct {
		desc   string
		input  int
		expOut uint
	}{
		{"Fact of zero", 0, 1},
		{"Fact of 1 ", 1, 1},
		{"Fact of neg number", -3, 0},
		{"Fact of pos number", 4, 24},
		{"Unsigned int overflow", 66, 0},
	}

	for i, tc := range cases {
		output := FactIterative(tc.input)

		if output != tc.expOut {
			t.Errorf("TEST[%d],failed.\nIterative: %s\nExpected:%v\nGot:%v\n", i, tc.desc, tc.expOut, output)
		}

		output = FactRecursive(tc.input)

		if output != tc.expOut {
			t.Errorf("TEST[%d],failed.\nRecursive: %s\nExpected:%v\nGot:%v\n", i, tc.desc, tc.expOut, output)
		}
	}
}

/*
	benchmark for Iterative approach
*/

func BenchmarkFactIterative(b *testing.B) {
	benchCases := []struct {
		desc  string
		input int
	}{
		{"factorial for 0", 0},
		{"factorial for 1", 1},
		{"factorial of 13", 13},
		{"factorial of 6", 6},
	}

	for _, v := range benchCases {
		for i := 0; i < b.N; i++ {
			FactIterative(v.input)
		}
	}
}

/*
	benchmark for recursive approach
*/

func BenchmarkFactRecursive(b *testing.B) {
	benchCases := []struct {
		desc  string
		input int
	}{
		{"factorial for 0", 0},
		{"factorial for 1", 1},
		{"factorial of 13", 13},
		{"factorial of 6", 6},
	}

	for _, v := range benchCases {
		for i := 0; i < b.N; i++ {
			FactRecursive(v.input)
		}
	}
}
