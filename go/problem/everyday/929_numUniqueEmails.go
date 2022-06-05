package everyday

import (
	"fmt"
	"strings"
)

func NumUniqueEmails(emails []string) int {
    /*
        根据每个邮件名，获得其对应的发送邮件名，存入hash 数组中
        最后统计数组长度即可
    */

    hash := make(map[string]int)

    for _, e := range emails {
        tmpEmail := strings.Split(e, "@") 

		localName := ""
		for _, v := range tmpEmail[0] {
			if v != '+' {
				if v != '.' {
					localName += string(v)
				}
			} else {
				break
			}
		}
		fmt.Println("old: ", e, "new: ", localName+"@"+tmpEmail[1])

		hash[localName+"@"+tmpEmail[1]] = 1
    }

	return len(hash)
}