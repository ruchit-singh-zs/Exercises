package prime

/*
	create a function to check for prime number
	Args : number of type int
	returns : string as an output
*/

func calculate(num int) bool {
	//Set flag default as false
	flag := false

	//if input < 1 : It's composite
	if num == 0 || num == 1 {
		return false
	}
	if num > 1 {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = true
				break
			}
		}
	}
	if flag == true {
		return false
	}
	return true
}
