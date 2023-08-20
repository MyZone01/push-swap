package pushswap

func isVoid(stack []int) bool {
	for _, v := range stack {
		if v != 0 {
			return false
		}
	}
	return true
}