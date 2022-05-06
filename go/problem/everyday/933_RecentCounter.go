package everyday

type RecentCounter struct {
	queue []int
}


func Constructor() RecentCounter {
	return RecentCounter{}
}


func (this *RecentCounter) Ping(t int) int {
	// 添加一个新请求
	this.queue = append(this.queue, t)
	// 将不在t-3000的请求，抛弃掉
	for this.queue[0] < t-3000 {
		this.queue = this.queue[1:]
	}
	return len(this.queue)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */