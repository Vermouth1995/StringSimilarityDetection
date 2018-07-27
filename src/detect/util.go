package detect

func MaxLen(str string, anotherStr string) int {
	if len(str) > len(anotherStr) {
		return len(str)
	}
	return len(anotherStr)
}
