package detect

func SimilarityDetect(a string, b string) float64 {
	distance := EditDistanceRecursion(a, b, len(a), len(b))
	ratio := float64(distance) / float64(MaxLen(a, b))
	return 1 - ratio
}

func EditDistanceRecursion(a string, b string, len_a int, len_b int) int {
	if len_a == 0 {
		return len_b
	} else if len_b == 0 {
		return len_a
	} else if string([]rune(a)[len_a-1:len_a]) == string([]rune(b)[len_b-1:len_b]) {
		return EditDistanceRecursion(a, b, len_a-1, len_b-1)
	} else {
		return MinVal(	EditDistanceRecursion(a, b, len(a)-1, len(b)) + 1,
						EditDistanceRecursion(a, b, len(a), len(b)-1) + 1,
						EditDistanceRecursion(a, b, len(a)-1, len(b)-1) + 1)
	}
}

func EditDistanceDynamic(a string, b string) int {
	return 1
}
