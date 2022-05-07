package everyday


func minMutation(start string, end string, bank []string) int {
	/*
		方法一：单向bfs
		1）排除边界条件：start 或 end 空或 bank 为空或 end 不在 bank 中。
		2）bfs 的初始化工作：初始化步长，初始化 queue，将 start 入队列，用 visit 来标记已经访问过的点。
		3）进行 bfs：先将步长+1，然后确定每次 bfs 的长度 size，寻找目标基因，然后入队出队等操作。
		时间复杂度：O(C∗n)
		空间复杂度：O(n)O(n)
	*/
	// 排除边界问题
	if len(start) == 0 || len(end) == 0|| len(bank) == 0 {
		return -1
	}

	bankMap := make(map[string]int, len(bank))
	for _, v := range bank {
		bankMap[v] = 0
	}

	if _, ok := bankMap[end]; !ok {
		return -1
	}
	// bfs的初始化工作
	var step int = 0
	queue := []string{start}

	// 开始bfs
	for len(queue) > 0 {
		// 当前队列，步长+1
		step++
		tmpQueue := queue
		queue = nil
		// 遍历整个队列数据
		for _, cur := range tmpQueue {
			for i, v := range cur {
				for _, v2 := range "ACGT" {
					if v != v2 {
						// 逐个字符替换成其不同的其他三个字符，拼接新的字符串
						nxt := cur[:i] + string(v2) + cur[i+1:]
						// 只有新字符串在bank里面，才继续
						if _, ok := bankMap[nxt]; ok {
							// 获取到目标字串，返回当前步长即可
							if nxt == end {
								return step
							}
							// 否则，将当前字符从bank移除，并添加进队列，即更新start集合
							// 当bank为空了，说明所有的字符都已经尝试过了，没有路径可以到达目标
							delete(bankMap, nxt)
							queue = append(queue, nxt)
						}
					}
				}
			}
		}
	}
	return -1
}

var mutationMap = map[uint8][3]string{
	'A': [...]string{"T", "G", "C"},
	'C': [...]string{"T", "G", "A"},
	'T': [...]string{"A", "G", "C"},
	'G': [...]string{"T", "A", "C"},
}

func idxOf(str string, bank []string) int {
	for i, s := range bank {
		if s == str {
			return i
		}
	}
	return -1
}

func minMutation2(start string, end string, bank []string) int {
	if idxOf(end, bank) == -1 {
		return -1
	}
	if start == end {
		return 0
	}
	count := 0
	isUsed := make([]bool, len(bank))

	startQueue := []string{start}
	endQueue := []string{end}
	for len(startQueue) > 0 {
		count++
		l := len(startQueue)
		for _, curr := range startQueue {
			for i := 0; i < len(curr); i++ {
				for _, c := range mutationMap[curr[i]] {
					gene := curr[:i] + c + curr[i+1:]
					if idx := idxOf(gene, endQueue); idx != -1 {
						return count
					}
					if idx := idxOf(gene, bank); idx != -1 && !isUsed[idx] {
						isUsed[idx] = true
						startQueue = append(startQueue, gene)
					}
				}
			}
		}
		startQueue = startQueue[l:]
		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
		}
	}
	return -1
}
