package top100


/*
输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]

*/

/*************************** 快排 ************************/
func sortArray(nums []int) []int {
	quickSort2(nums, 0, len(nums)-1)
	return nums
}

func quickSort2(nums []int, left, right int) {
	if left < right {
		pivot := partition2(nums, left, right)

		quickSort2(nums, left, pivot-1)
		quickSort2(nums, pivot+1, right)
	}
}

func partition2(nums []int, left, right int) int {
	pivot := nums[left]

	for left < right {
		// 先从右边找到小于基准值的数
		for left < right && nums[right] >= pivot {
			right--
		}

		// 用这个数更新基准值位置
		nums[left] = nums[right]

		// 接下来从左边找大于基准值的数
		for left < right && nums[left] < pivot {
			left++
		}

		// 用这个数更新之前空出来的位置
		nums[right] = nums[left]
	}
	// 最后更新基准值到最“中间”的位置
	nums[left] = pivot
	return left
}

/*************************** 堆排 ************************/
func SortArray2(nums []int) []int {
	heapSort(nums, len(nums)-1)
	return nums
}

func heapSort(nums []int, len int) {
	// 初始化，i从最后一个父节点开始调整
	for i := len/2; i >= 0; i-- {
		maxHeapify(nums, i, len)
	}

	// 先将第一个元素和已经排好的元素的前一位做交换，再重新调整（刚调整的元素之前的元素），
	// 直到排序完成
	for i := len; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(nums, 0, i-1)
	}
}

func maxHeapify(nums []int, start, end int) {
	// 建立父节点和子节点
	dad := start
	son := dad * 2 + 1

	// 若子节点指标在范围内才做比较
	for son <= end {
		// 选择两个子节点较大的
		if son + 1 <= end && nums[son] < nums[son+1] {
			son++
		}
		// 如果父节点大于子节点代表调整完毕，直接跳出
		if nums[dad] > nums[son] {
			return
		} else {
			// 否则交换父子节点再继续子节点和孙节点比较
			nums[dad], nums[son] = nums[son], nums[dad]
			dad = son
			son = dad * 2 + 1
		}
	}
}