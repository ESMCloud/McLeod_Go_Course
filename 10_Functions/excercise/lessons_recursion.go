package excercise

//FactorialFunc1 --> dont use, use loops
func FactorialFunc1(x int) int {
	// if x == 0 {
	// 	return 1
	// }
	// return x * FactorialFunc1(x-1)

	for i := x - 1; i > 0; i-- {
		x *= (i)
	}
	return x
}
