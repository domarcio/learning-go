package math

/*
Divider is a simple function to divider arg1 * arg2 and attach to assign result variable
The remainder variable is a remainder of division
*/
func Divider(arg1 int, arg2 int) (result int, remainder int) {
	result = arg1 / arg2
	remainder = arg1 % arg2
	return
}
