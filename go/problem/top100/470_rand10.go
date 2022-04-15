package top100

//  把已知的范围按等概率划分，分别返回0和1。 然后通过0和1组合得到任意范围的结果。
func rand10() int {
	var res int

	for {
		res = get1Or0()<<3 + get1Or0()<<2 + get1Or0()<<1 + get1Or0()
		if res <= 9 {
			break
		}
	}

	return res + 1
}

func get1Or0() int {
	num := rand7()

	for {
		if num < 4 { // 50%
			return 0
		} else if num > 4 { // 50%
			return 1
		} else {
			num = rand7()
		}
	}
}

func rand7() int {
	return 0
}

func Rand10() int {
    /*
		1、先用一个通用的函数，取出等概率的0和1；
		2、然后位运算（0，1）取目标值；
	*/
	var res int

	for {
		/*
			[0, 15] -> [1, 10]
		*/
		res = getBinary() << 3 + getBinary() << 2 + getBinary() << 1 + getBinary() 
		if res <= 9 {
			break 
		}
	}

	return res + 1
}

func getBinary() int {
	num := rand7()

	for {
		if num < 4 {
			return 0
		} else if num > 4 {
			return 1
		} else {
			num = rand7()
		}
	}
}