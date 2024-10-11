package piscine

func Atoi(s string) int {
	numbers := []int{}
	isnegative := false
	if s[0] == '-' {
		numbers = append(numbers, 0)
		isnegative = true
		for i := 1; i <= len(s)-1; i++ {
			if s[i] >= '0' && s[i] <= '9' {
				numbers = append(numbers, int(s[i]-'0'))
			} else {
				return 0
			}
		}
	} else if s[0] == '+' {
		numbers = append(numbers, 0)
		for i := 1; i <= len(s)-1; i++ {
			if s[i] >= '0' && s[i] <= '9' {
				numbers = append(numbers, int(s[i]-'0'))
			} else {
				return 0
			}
		}
	} else {
		for i := 0; i <= len(s)-1; i++ {
			if s[i] >= '0' && s[i] <= '9' {
				numbers = append(numbers, int(s[i]-'0'))
			} else {
				return 0
			}
		}
	}
	result := 0
	for _, num := range numbers {
		result = result*10 + num
	}
	if isnegative {
		return -result
	}
	return result
}
