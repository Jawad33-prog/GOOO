package piscine

func Capitalize(s string) string {
	capitalizeNext := true
	result := []byte(s)

	for i := 0; i < len(result); i++ {
		if (result[i] >= 'A' && result[i] <= 'Z') || (result[i] >= 'a' && result[i] <= 'z') || (result[i] >= '0' && result[i] <= '9') {
			if capitalizeNext && result[i] >= 'a' && result[i] <= 'z' {
				result[i] -= 'a' - 'A'
			} else if !capitalizeNext && result[i] >= 'A' && result[i] <= 'Z' {
				result[i] += 'a' - 'A'
			}
			capitalizeNext = false
		} else {
			capitalizeNext = true
		}
	}

	return string(result)
}
