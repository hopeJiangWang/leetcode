package common

func MyMinInt(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func MyMaxInt(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func MyAbsInt64(tm int64) int64 {
	y := tm >> 63
	return (tm ^ y) - y
}

func MyAbsInt(tm int) int {
	y := tm >> 31
	return (tm ^ y) - y
}