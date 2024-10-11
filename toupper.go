package piscine

func ToUpper(s string) string {
	var result []rune
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			result = append(result, r-32)
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
