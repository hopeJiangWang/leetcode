package problem

// import "encoding/binary"

// ["abcw","baz","foo","bar","xtfn","abcdef"]
// "abcw", "xtfn" -> 16

func MaxProduct1(words []string) int {
    res := 0
    for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			// 有相同的字符，则跳过
			if stringToBinary(words[i]) & stringToBinary(words[j]) != 0 {
				continue
			}

			tmp := len(words[i]) * len(words[j])
			if tmp > res {
				res = tmp
			}
		}
	}

    return res
}

// 将字符串转换成二进制数
func stringToBinary(s string) int {
	res := 0
	for _, v := range s {
		res |= (1 << (v - 'a'))
	}

	return res
}


func MaxProduct(words []string) int {
    res := 0
	wordsLen := len(words)

	// 先预处理，减少循环
	binaryWords := make([]int, wordsLen)
	for i := 0; i < wordsLen; i++ {
		for _, v := range words[i] {
			binaryWords[i] |= (1 << (v - 'a'))
		}
	}

    for i := 0; i < wordsLen; i++ {
		for j := i + 1; j < wordsLen; j++ {
			// 有相同的字符，则跳过
			if binaryWords[i] & binaryWords[j] != 0 {
				continue
			}

			tmp := len(words[i]) * len(words[j])
			if tmp > res {
				res = tmp
			}
		}
	}

    return res
}
