package twofer

import (
	"fmt"
	"testing"
)

/*
create a function to check for Table Driven Test
	Args : pointer to testing package type T
	checks whether the expected output is predicted output or not
*/

func TestTwoFer(t *testing.T) {

	fmt.Println("** Table Driven Test **")
	fmt.Println("-----------------------")

	//Construct a struct
	cases := []struct {
		input           string
		predictedOutput string
	}{
		{"ruchit", "One for ruchit,one for me."},
		{"name", "One for you, one for me."},
		{"aryan", "One for aryan,one for me."},
	}

	//compute and test
	for i := 0; i < len(cases); i++ {
		receivedOutput := compute(cases[i].input)
		if receivedOutput != cases[i].predictedOutput {
			t.Errorf("got : %v :: expected : %v", receivedOutput, cases[i].predictedOutput)
		}
	}
}
