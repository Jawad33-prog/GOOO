package piscine

func NRune(s string, n int) rune {
	if n > len(s) || n <= 0 {
		return 0
	}
	nthrune := []rune(s)[n-1]
	return nthrune
}
