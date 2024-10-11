package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		printNV()
		return
	}

	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		if nbr == -9223372036854775808 {
			nbr = -9223372036854775807 - 1
		} else {
			nbr = -nbr
		}
	}

	baseLen := len(base)
	result := ""
	for nbr > 0 {
		result = string(base[nbr%baseLen]) + result
		nbr /= baseLen
	}

	for _, char := range result {
		z01.PrintRune(char)
	}
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' || seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func printNV() {
	z01.PrintRune('N')
	z01.PrintRune('V')
}
