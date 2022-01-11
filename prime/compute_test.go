package prime

import (
	"fmt"
	"testing"
)

/*
	create a function to check for Table Driven Test
	Args : pointer to testing package type T
	checks whether the expected output is predicted output or not
*/

func TestPrime(t *testing.T) {

	fmt.Println("** Table Driven Test **")
	fmt.Println("-----------------------")

	//Construct a struct entity with fields as number and output

	cases := []struct {
		input           int
		predictedOutput bool
	}{
		{10, false},
		{7, true},
		{1, false},
	}

	//compute and test
	for i := 0; i < len(cases); i++ {
		receivedOutput := calculate(cases[i].input)

		if receivedOutput != cases[i].predictedOutput {
			t.Errorf("Test Case Failed [%d] \n got %v, expected %v", i, receivedOutput, cases[i].predictedOutput)
		}
	}
}
