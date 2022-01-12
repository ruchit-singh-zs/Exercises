package factorial

func FactIterative(num int) uint {
	var fact uint = 1

	if num < 0 {
		return 0
	}

	if num == 0 {
		return 1
	}

	for i := 1; i <= num; i++ {
		fact = fact * uint(i)
	}
	return fact
}

/*
    compute Factorial in recursive manner
	Args : input number
	returns: Factorial
*/

func FactRecursive(num int) uint {
	if num < 0 {
		return 0
	}

	if num == 0 || num == 1 {
		return 1
	}

	return uint(num) * FactRecursive(num-1)
}
