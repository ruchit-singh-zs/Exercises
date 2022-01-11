package triangle

/*
	create function to compute the type of triangle
	Args: 3 sides of type int each
	returns string as an output
*/

func compute(x, y, z int) string {
	if x+y <= z || x+z <= y || y+z <= x {
		return "Degenerate Triangle"
	}
	if x == y && y == z && x == z {
		return "Equilateral Triangle"
	}
	if x != y && y != z && x != z {
		return "Scalene Triangle"
	}
	return "Isosceles Triangle"
}
