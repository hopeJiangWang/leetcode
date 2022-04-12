package top100

/*
输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
*/

func RemoveDuplicates([]int nums) int {
	/*
		使用双指针，左指针计数用，右指针跳过相同数据
	*/

	left, right, lenOfNums := 0, 0, len(nums)
	for right < lenOfNums {
		if right < lenOfNums && nums[left] == nums[right] {
			right++
            // fmt.Println("left: ", left, "right: ", right, "nums: ", nums)
		} else if right < lenOfNums {
            left++
			nums[left] = nums[right]
            // fmt.Println("left: ", left, "right: ", right, "nums: ", nums)
		}
	}

	return left+1	// 需要加上right位置
}
