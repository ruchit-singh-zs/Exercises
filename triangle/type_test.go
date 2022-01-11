package triangle

import (
	"fmt"
	"testing"
)

/*
create a function to check for Table Driven Test
	Args : pointer to testing package type T
	checks whether the expected output is predicted output or not
*/

func TestTriangle(t *testing.T) {

	fmt.Println("** Table Driven Test **")
	fmt.Println("-----------------------")

	//Construct a struct
	cases := []struct {
		side1, side2, side3 int
		predictedOutput     string
	}{
		{3, 3, 3, "Equilateral Triangle"},
		{4, 5, 6, "Scalene Triangle"},
		{5, 5, 6, "Isosceles Triangle"},
		{1, 2, 3, "Degenerate Triangle"},
	}

	//compute and test
	for i := 0; i < len(cases); i++ {
		receivedOutput := compute(cases[i].side1, cases[i].side2, cases[i].side3)

		if receivedOutput != cases[i].predictedOutput {
			t.Errorf("Test[%d] failed \n got : %v :: expected : %v", i, receivedOutput, cases[i].predictedOutput)
		}
	}
}
