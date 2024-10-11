package piscine

func IterativePower(nb int, power int) int {
	multi := nb
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if nb == 1 {
		return 1
	} else if power == 1 {
		return nb
	} else {
		for i := 0; i < power-1; i++ {
			multi = multi * nb
		}
	}
	return multi
}
