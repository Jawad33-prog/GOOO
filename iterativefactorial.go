package piscine

func IterativeFactorial(nb int) int {
	factorial := 1
	if nb < 0 || nb > 20 {
		return 0
	} else {
		for i := 2; i <= nb; i++ {
			factorial = factorial * i
		}
	}
	return factorial
}
