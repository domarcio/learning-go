package math

/*
Calculate is a generic function to calculate arg1 and arg1 based on function
*/
func Calculate(function func(int, int) int, arg1 int, arg2 int) int {
	return function(arg1, arg2)
}
