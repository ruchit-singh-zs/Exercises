package bob

import (
	"fmt"
	"testing"
)

/*
create a function to check for Table Driven Test
	Args : pointer to testing package type T
	checks whether the expected output is predicted output or not
*/

func TestBob(t *testing.T) {

	fmt.Println("** Table Driven Test **")
	fmt.Println("-----------------------")

	//Construct a struct entity with fields as number and output

	cases := []struct {
		input          string
		expectedOutput string
	}{
		{"random", "Whatever"},
		{"Hi?", "Sure"},
		{"HI", "Whoa, chill out!"},
		{"HI?", "Calm down, I know what I'm doing!"},
	}

	//compute and test
	for i := 0; i < len(cases); i++ {
		receivedOutput := respond(cases[i].input)

		if receivedOutput != cases[i].expectedOutput {
			t.Errorf("Test Failed[%d] \n Got %v :: Expected %v", i, receivedOutput, cases[i].expectedOutput)
		}
	}
}
