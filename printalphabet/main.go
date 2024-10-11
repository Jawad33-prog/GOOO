package main

import "github.com/01-edu/z01"

func main() {
	for lettre := 'a'; lettre <= 'z'; lettre++ {
		z01.PrintRune(lettre)
	}
	z01.PrintRune('\n')
}
