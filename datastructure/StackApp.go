package TestProject


func fabi(i int) int {
	if i < 2 {
		if i == 0 {
			return 0
		}
		return 1
	}
	return fabi(i-1) + fabi(i-2)
}

func Convert2PostfixExp(stack Stack,midExp string) string{
	return ""
}