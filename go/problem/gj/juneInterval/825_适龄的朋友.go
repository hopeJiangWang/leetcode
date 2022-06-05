package juneinterval

import "sort"

func NumFriendRequests(ages []int) int {
	/*
		也即查找统计: 0.5*ages[x]+7 < ages[y] <= ages[x]
		可以得到以下几点：
		1、ages[y]需要大于14岁；
		2、对于当前ages[x]，如果当前左侧的年龄小于左侧标准，left++；
		3、当前右侧年龄不大于右侧标准，right++
		4、左右指针之差，就是结果
		
		从大到小排序之后，枚举每个ages[x]，统计其中的ages[y]
	*/
	sort.Ints(ages)

	len1, res := len(ages), 0
	left, right := 0, 0
	for _, age := range ages {
        if age < 15 {
            continue
        }

		for ages[left] <= age/2+7 {
			left++
		}

		for right+1 < len1 && ages[right+1] <= age {
			right++
		}

		res += right - left
	}
	return res
}
