package cal

// AddUpper 1+2+3+...
func AddUpper(num int) int {
	res := 0
	for i := 1; i <= num; i++ {
		res += i
	}
	return res
}
