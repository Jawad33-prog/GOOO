package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	result := RecursiveFactorial(nb - 1)
	return result * nb
}
