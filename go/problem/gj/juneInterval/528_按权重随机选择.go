package juneinterval

import (
	"crypto/rand"
	"sort"
)

/*
w = [3,1,2,4], total = 10
按照[1,3],[4,4],[5,6],[7,10]划分，使得其长度恰好依次是3，1，2，4
pre[i]表示w的前缀和

那么第i个区间的左边界就是它之前出现的所有元素的和加上 1:
pre[i]-w[i]+1

右边界就是到它为止的所有元素之和:
pre[i]

那么我们希望找到满足：
pre[i]-w[i]+1 <= x <= pre[i]

pre[i]递增，使用二分查找即可

*/

type Solution struct {
	pre []int
}


func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{w}
}


func (this *Solution) PickIndex() int {
	x := rand.Intn(this.pre[len(this.pre)-1]) + 1
	return sort.SearchInts(this.pre, x)
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */