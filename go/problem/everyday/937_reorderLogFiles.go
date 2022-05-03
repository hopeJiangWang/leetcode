package everyday

import (
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	/*
		1、字母日志要在数字日志之前：let > dig
		2、字母日志按内容字母顺序排序
		3、数字日志保持原有相对排序
	*/

	sort.SliceStable(logs, func(i, j int) bool {
		// 按空格拆分成两部分，这里只需要取第二部分即可
		slice1 := strings.SplitN(logs[i], " ", 2)[1]
		slice2 := strings.SplitN(logs[j], " ", 2)[1]
		// 取出第二部分的第一个字符判断类型
		isDigit1 := unicode.IsDigit(rune(slice1[0]))
		isDigit2 := unicode.IsDigit(rune(slice2[0]))
		// 都是数字字符，直接按原有向度排序即可
		if isDigit1 && isDigit2 {
			return false
		}
		// 都是字母字符，按字母顺序排即可
		if !isDigit1 && !isDigit2 {
			return slice1 < slice2 || slice1 == slice2 && logs[i] < logs[j]
		}
		// 否则，就是一个是数字一个不是，那就看第一个是不是数字即可
		return !isDigit1
	})

	return logs
}