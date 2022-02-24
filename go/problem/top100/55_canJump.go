package top100

/*
输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。

输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
*/

func CanJump(nums []int) bool {
	/*
		贪心：一直往后面走，计算出能到达的最长的距离，和数组长度做一个比较
		如果能跳到最后，就成功了
	*/
	lenNums := len(nums)
	var maxRoute int

	for i := 0; i < lenNums; i++ {
		// 如果能到达的最远距离都比当前位置更短，那肯定到达不了
		if i > maxRoute {
			return false
		}
		maxRoute = max(maxRoute, i+nums[i]) // 当前位置能到达的最远距离
	}

	return true
}
