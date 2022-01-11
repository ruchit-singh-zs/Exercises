package leap_year

/*
	create a function to check for leap year
	Args : year of type int
	returns : string as an output
*/

func compute(year int) bool {
	if year%4 == 0 {
		if year%100 != 0 || year%400 == 0 {
			return true
		}
	}
	return false
}
