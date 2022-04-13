package top100

/*
输入：[2,2,1,1,1,2,2]
输出：2
*/

func MajorityElement(nums []int) int {
	/*
		摩尔投票法（依次去除两个不同的数，那么剩下来的就是众数了）：
		1. 先选择一个 candidate, 初始化计数count=0, 然后遍历数组;
		2. 如果x == candidate, 那么count++; 如果x != candidate, 那么count--;
		3. 如果count == 0, 这时候更新candidate为x;
	*/
	candidate := nums[0]
	count := 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
        
        if v == candidate {
			count ++
		} else {
			count --
		}
	}

	return candidate
}