package detect

func MaxLen(str string, anotherStr string) int {
	if len(str) > len(anotherStr) {
		return len(str)
	}
	return len(anotherStr)
}

func MinVal(first int, args... int) int {
	for _ , v := range args{
		if first > v {
			first = v
		}
	}
	return first
}
