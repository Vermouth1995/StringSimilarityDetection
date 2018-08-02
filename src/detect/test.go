package detect

import (
	"fmt"
)

func SimilarityDetectionTest() {
	var i int
	var testStr = [3][2]string{
		{"", "ebcg"},
		{"ab", "ab"},
		{"abcdefg", "bcdefg"}}
	for i = 0; i < 3; i++ {
		fmt.Printf("%s 和 %s 的相似度: %f\n", testStr[i][0], testStr[i][1], SimilarityDetect(testStr[i][0], testStr[i][1]))
	}
}
