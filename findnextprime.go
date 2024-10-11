package piscine

func FindNextPrime(n int) int {
	if n <= 1 {
		return 2
	}
	if n == 2 {
		return 2
	}
	if n%2 == 0 {
		n++
	}
	for {
		if n <= 3 {
			return n
		}
		if n%2 != 0 && n%3 != 0 {
			isPrime := true
			for i := 5; i*i <= n; i += 6 {
				if n%i == 0 || n%(i+2) == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				return n
			}
		}
		n += 2
	}
}
