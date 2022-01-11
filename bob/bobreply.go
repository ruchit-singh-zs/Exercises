package bob

import "strings"

/*
	create a function compute()
	Args: question of type string
	returns output of type string
*/

func respond(input string) string {
	switch {
	case input == strings.ToUpper(input) && strings.HasSuffix(input, "?"):
		return "Calm down, I know what I'm doing!"
	case input == strings.ToUpper(input):
		return "Whoa, chill out!"
	case strings.HasSuffix(input, "?"):
		return "Sure"
	case input == "":
		return "Fine. Be that way!"
	default:
		return "Whatever"
	}
}
