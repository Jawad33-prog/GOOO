package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	for i := 2; i <= nb; i++ {
		if nb == i*i {
			return i
		}
	}
	return 0
}
