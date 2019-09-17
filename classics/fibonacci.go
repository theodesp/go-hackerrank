package classics

func loop(a, b, n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return a
	}
	return loop(b, a + b, n-1)
}

func GetNthFib(n int) int {
	return loop(0, 1, n)
}



