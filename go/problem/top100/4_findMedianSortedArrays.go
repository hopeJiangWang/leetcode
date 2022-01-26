package top100

import (
	"fmt"
)

/*
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000

输入：nums1 = [], nums2 = [1]
输出：1.00000

输入：nums1 = [2], nums2 = []
输出：2.00000
*/

func FindMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	/*
		1: 直接合并
		2: 可以知道总的长度，就可以转换成找第k大的数这种
	*/
	// sumLen := len(nums1) + len(nums2)
	var res []int

	l1, l2 := 0, 0

	for l1 < len(nums1) && l2 < len(nums2) {
		if nums1[l1] <= nums2[l2] {
			res = append(res, nums1[l1])
			l1++
		} else if l2 < len(nums2) {
			res = append(res, nums2[l2])
			l2++
		}
	}

	res = append(res, nums1[l1:]...)
	res = append(res, nums2[l2:]...)

	fmt.Printf("res: %v", res)

	if len(res) % 2 == 0 {
		return float64(res[len(res)/2] + res[len(res)/2-1]) / 2.0
	} else {
		return float64(res[len(res)/2])
	}
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	/*
		1: 直接合并
		2: 可以知道总的长度，就可以转换成找第k大的数这种
	*/
	sumLen := len(nums1) + len(nums2)

	if sumLen % 2 == 0 {
		return float64(FindKthNumber(nums1, nums2, sumLen/2+1) + FindKthNumber(nums1, nums2, sumLen/2)) / 2.0
	} else {
		return float64(FindKthNumber(nums1, nums2, sumLen/2+1))
	}
}

func FindKthNumber(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0

	for {
		// 只剩下一个数组有数据了，就直接找第k个就可以了
		if index1 == len(nums1) {
            return nums2[index2 + k - 1]
        }
        if index2 == len(nums2) {
            return nums1[index1 + k - 1]
        }

		// 或者 k 为 1 了，那我们直接找最前面的两个数中的最小值即可
        if k == 1 {
            return min(nums1[index1], nums2[index2])
        }

		half := k/2
		// 根据k，取每个数组的k/2-1的值
        newIndex1 := min(index1 + half, len(nums1)) - 1
        newIndex2 := min(index2 + half, len(nums2)) - 1
        pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]

		// 根据情况，更新k值
		if pivot1 <= pivot2 {
            k -= (newIndex1 - index1 + 1)
            index1 = newIndex1 + 1
        } else {
            k -= (newIndex2 - index2 + 1)
            index2 = newIndex2 + 1
        }
	}
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}