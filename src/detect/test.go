package detect

import (
	"fmt"
)

func SimilarityDetectionTest() {
	var i int
	var testStr = [4][2]string{
		{"", "ebcg"},
		{"ab", "ab"},
		{"abcdefg", "bcdefg"},
		{"abcdefg", "qeqwebcde"}}
	for i = 0; i < len(testStr); i++ {
		fmt.Printf("%s 和 %s 的相似度: %f\n", testStr[i][0], testStr[i][1], SimilarityDetect(testStr[i][0], testStr[i][1]))
	}
}
