package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var digits []rune
	for n > 0 {
		digit := rune('0' + n%10)
		digits = append([]rune{digit}, digits...)
		n /= 10
	}
	for _, r := range digits {
		z01.PrintRune(r)
	}
}
