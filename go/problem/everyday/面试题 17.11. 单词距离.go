package everyday

/*
输入：words = ["I","am","a","student","from","a","university","student","in","a","city"], word1 = "a", word2 = "student"
输出：1

*/

func FindClosest(words []string, word1 string, word2 string) int {
	var res int = 100001

	/*
		arr_a: [2,5,9]
		arr_s: [3,7]
		即找这两个数组直接任意两个值之间的最小差值

		可以不记录数组，直接处理：
		因为都是从头开始处理，只需要比较最新的两个位置之间的距离，取最小值即可
	*/
	wordsLen := len(words)
	first, second := 0, 0
	for i := 0; i < wordsLen; i++ {
		if words[i] == word1 {
			first = i
		} 

		if words[i] == word2 {
			second = i
		}

		if first > 0 && second > 0 {
			res = min(res, MyAbsInt(first-second))
		}
	}

	return res
}

// func MyAbsInt(tm int) int {
// 	y := tm >> 31
// 	return (tm ^ y) - y
// }

// func min(a, b int) int {
//     if a > b {
//         return b
//     }
//     return a
// }
