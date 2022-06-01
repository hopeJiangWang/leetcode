package juneinterval

type NumArray struct {
	Nums []int 
}


func Constructor(nums []int) NumArray {
	return NumArray {
		Nums: nums,
	}
}

// 求区间和
func (this *NumArray) SumRange(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += this.Nums[i]
	}

	return sum
}
