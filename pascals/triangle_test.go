package pascal

import (
	"testing"
)

// TestPascal test if the values of pascal triangle are correct or not
func TestPascal(t *testing.T) {
	expOut := `           1
          1 1
         1 2 1
        1 3 3 1
       1 4 6 4 1
      1 5 10 10 5 1
     1 6 15 20 15 6 1
    1 7 21 35 35 21 7 1
   1 8 28 56 70 56 28 8 1
  1 9 36 84 126 126 84 36 9 1
`

	cases := []struct {
		desc   string
		input  int
		output string
	}{
		{"zero row", 0, ""},
		{"ten row", 10, expOut},
	}

	for i, tc := range cases {
		output := drawTriangle(tc.input)

		// Check if the output is correct or not
		// if not correct return the error using Errorf function of testing package
		if output != tc.output {
			t.Errorf("TEST[%d],failed. %s\nExpected : %v \nGot : %v", i, tc.desc, tc.output, output)
		}
	}
}
