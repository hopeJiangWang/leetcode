package top100

// import "fmt"

/*
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6

输入：height = [4,2,0,3,2,5]
输出：9
*/

func Trap1(height []int) int {
	var res int = 0
	/*
		从左至右遍历，依次获取当前位置 i 的左右两边的最大值 max_l, max_r, h_min = min(max_l, max_r)
		那么本位置可以盛水的量为 h_min-height[i](小于0表示无法盛水)
	*/
	if len(height) == 0 {
		return 0
	}

	for i := 0; i < len(height); i++ {
		maxL, maxR := findLRMax(height, i)
		if min(maxL, maxR)-height[i] > 0 {
			res += min(maxL, maxR) - height[i]
		}
	}

	return res
}

func findLRMax(arr []int, index int) (int, int) {
	var maxL int = 0
	var maxR int = 0
	for i := 0; i < len(arr); i++ {
		if i < index {
			maxL = max(maxL, arr[i])
		} else if i > index {
			maxR = max(maxR, arr[i])
		}
	}
	return maxL, maxR
}

func Trap2(height []int) int {
	var res int = 0
	n := len(height)
	/*
		从左至右遍历，依次获取当前位置 i 的左右两边的最大值 max_l, max_r, h_min = min(max_l, max_r)
		那么本位置可以盛水的量为 h_min-height[i](小于0表示无法盛水)

		优化一下，先遍历，找出每个位置对应的h_min，存起来，这样子就从O(n^2) -> O(n)
	*/
	if n == 0 {
		return 0
	}

	maxL := make([]int, n)
	maxR := make([]int, n)
	for i := 1; i < n; i++ {
		maxL[i] = max(maxL[i-1], height[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		maxR[i] = max(maxR[i+1], height[i+1])
	}

	for i := 0; i < n; i++ {
		if min(maxL[i], maxR[i])-height[i] > 0 {
			res += min(maxL[i], maxR[i]) - height[i]
		}
	}

	return res
}

func Trap(height []int) int {
	var res int = 0
	n := len(height) // 只计算一次也能节省很多时间
	/*
		从左至右遍历，依次获取当前位置 i 的左右两边的最大值 max_l, max_r, h_min = min(max_l, max_r)
		那么本位置可以盛水的量为 h_min-height[i](小于0表示无法盛水)

		优化一下，先遍历，找出每个位置对应的h_min，存起来，这样子就从O(n^2) -> O(n)

		上面的优化之后，空间复杂度就上来了，考虑用双指针来降低空间复杂度：
		left=0, right=n-1, leftMax=0, rightMax=0
		两个指针从两端向中间移动：
			1.height[left] >= height[right]: 此时 leftMax >= rightMax, right 位置可以盛水：
			rightMax − height[right], right++
			2.height[left] < height[right]: 此时 leftMax < rightMax, left 位置可以盛水：
			leftMax − height[left], left++
	*/
	if n == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}
	return res
}
