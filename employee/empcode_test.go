package employee

import (
	"testing"
)

/*
	create a test function to see the test driven development
*/

func TestCheckAge(t *testing.T) {
	emp := Employee{"age is greater than 22", 24, "aryan", "2"}
	cases := []struct {
		desc           string
		age            int
		name           string
		id             string
		expectedOutput Employee
		ok             bool
	}{
		{"age is less than 22", 20, "ruchit", "1", Employee{}, false},
		{"age is greater than 22", 24, "aryan", "2", emp, true},
	}

	for i, tc := range cases {
		actualOutput, actualBool := CheckAge(tc.desc, tc.age, tc.name, tc.id)

		if actualOutput != cases[i].expectedOutput {
			t.Errorf("Test Failed[%d] \n Got %v :: Expected %v", i, actualOutput, cases[i].expectedOutput)
		}

		if actualBool != cases[i].ok {
			t.Errorf("Test Failed[%d] \n Got %v :: Expected %v", i, actualBool, cases[i].ok)
		}
	}
}
