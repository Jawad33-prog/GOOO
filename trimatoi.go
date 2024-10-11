package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	started := false

	for _, char := range s {
		if char == '-' && !started {
			sign = -1
			started = true
		} else if char >= '0' && char <= '9' {
			started = true
			result = result*10 + int(char-'0')
		}
	}

	return sign * result
}
