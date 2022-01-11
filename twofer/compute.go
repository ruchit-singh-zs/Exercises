package twofer

/*
	create a function compute to perform the operation
	Args : user input of type string
	returns an output of type string
*/

func compute(input string) string {

	if input == "name" {
		return "One for you, one for me."
	}
	return "One for " + input + ",one for me."
}
