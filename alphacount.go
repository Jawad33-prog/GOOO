package piscine

func AlphaCount(s string) int {
	lettercount := 0
	for i := 0; i <= len(s)-1; i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			lettercount += 1
		}
	}
	return lettercount
}
