package detect

func SimilarityDetect(a string, b string) float64 {
	distance := EditDistanceDynamic(a, b)
	ratio := float64(distance) / float64(MaxLen(a, b))
	return 1 - ratio
}

/**
 * 递归算法
 * @param {string} a
 * @param {string} b
 * @param {int} len_a 字符串 a 的长度
 * @param {int} len_b 字符串 b 的长度
 * @return {int} 从 a → b 的最小编辑距离
 */
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

/**
 * 动态规划算法
 * @param {string} a
 * @param {string} b
 * @return {int} 从 a → b 的最小编辑距离
 */
func EditDistanceDynamic(a string, b string) int {
	var (
		len_a int = len(a)
		len_b int = len(b)
		i int
		j int
	)

	if len_a == 0 {
		return len_b
	}
	if len_b == 0 {
		return len_a
	}

	dis := make([][]int, len_a)
    for i = 0; i < len_a; i++ {
        dis[i] = make([]int, len_b)
    }

	for i = 0; i < len_a; i++ {
        dis[i][0] = i;
    }
    for j = 0; j < len_b; j++ {
        dis[0][j] = j;
    }

	for i = 1; i < len_a; i++ {
        for j = 1; j < len_b; j++ {
            if a[i-1] == b[j-1] {
                dis[i][j] = dis[i-1][j-1];
            } else {
                dis[i][j] = MinVal(dis[i-1][j]+1, dis[i][j-1]+1, dis[i-1][j-1]+1);
            }
        }
    }

	return dis[len_a-1][len_b-1]
}
