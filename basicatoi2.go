package piscine

func BasicAtoi2(s string) int {
	numbers := []int{}
	for i := 0; i <= len(s)-1; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			numbers = append(numbers, int(s[i]-'0'))
		} else {
			return 0
		}
	}
	result := 0
	for _, num := range numbers {
		result = result*10 + num
	}
	return result
}
