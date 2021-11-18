package problem

import (
	"math"
	// "mycode/leetcode/go/problem/common"
)

func BulbSwitch(n int) int {
	/*
		法1. 使用一个bool数组，记录每盏灯的状态，并依次进行操作，最后统计；
		法2. 直接平方根：
			1. 一个编号为 x 的灯泡经过 n 轮后处于打开状态的充要条件为「该灯泡被切换状态次数为奇数次」。
			同时，一个灯泡切换状态的次数为其约数的个数（去重）。
			于是问题转换为：在 [1,n] 内有多少个数，其约数的个数为奇数。这些约数个数为奇数的灯泡就是最后亮着的灯泡。
			又根据「约数」的定义，我们知道如果某个数 k 为 x 的约数，那么 x/k 亦为 x 的约数, 
			即「约数」总是成对出现，那么某个数的约数个数为奇数，意味着某个约数在分解过程中出现了 2 次，
			且必然重复出现在同一次拆解中，即 k = x/k, k= x，即有 xx 为完全平方数（反之亦然）。
			问题最终转换为：在 [1,n] 中完全平方数的个数为多少。
			根据数论推论，[1,n] 中完全平方数的个数为 sqrt{n} ，即最后亮着的灯泡数量为 sqrt{n} 
	*/

	return int(math.Sqrt(float64(n) + 0.5))
}