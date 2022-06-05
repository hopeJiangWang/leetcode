package every202206

import "math/rand"

type Solution struct {
	Radius, Xcenter, Ycenter float64
}

/*
拒绝采样方法：
	在一个更大的范围内生成随机数，然后拒绝掉那些不在题目范围内的数据
	这里采用一个边长为2R的正方形

*/
func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{radius, x_center, y_center}
}

func (this *Solution) RandPoint() []float64 {
	for {
		x := rand.Float64()*2 - 1	// [-1, 1)
		y := rand.Float64()*2 - 1

		// 在圆内，就返回他
		if x*x + y*y < 1 {
			return []float64{this.Xcenter + this.Radius*x, this.Ycenter + this.Radius*y}
		}
	}
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */

