package factorial

import "testing"

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

func BenchmarkFactorialIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactIterative(17)
	}
}
func BenchmarkFactorialRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactRecursive(17)
	}
}
