package leap_year

import (
	"fmt"
	"testing"
)

/*
create a function to check for Table Driven Test
	Args : pointer to testing package type T
	checks whether the expected output is predicted output or not
*/

func TestYear(t *testing.T) {
	fmt.Println("** Table Driven Test **")
	fmt.Println("-----------------------")

	//Construct a struct
	cases := []struct {
		year   int
		result bool
	}{
		{2000, true},
		{1987, false},
	}

	//compute and test
	for i := range cases {
		receivedOutput := compute(cases[i].year)

		if receivedOutput != cases[i].result {
			t.Errorf("Test Case Failed [%d] \n got %v :: expected %v", i, receivedOutput, cases[i].result)
		}
	}
}
