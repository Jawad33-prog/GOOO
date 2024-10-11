package piscine

func StrRev(s string) string {
	val := []rune(s)
	for a, b := 0, len(val)-1; a < b; a, b = a+1, b-1 {
		val[a], val[b] = val[b], val[a]
	}
	return string(val)
}
