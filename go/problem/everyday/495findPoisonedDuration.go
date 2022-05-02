package everyday

func findPoisonedDuration(timeSeries []int, duration int) int {
    // 1. 如果没中毒：expired = timeSeries[i] + duration，增加的持续中毒时间为duration
    // 2. 如果已中毒：增加的持续中毒时间为timeSeries[i] + duration - expired
    expired, res := 0, 0
    for _, v := range timeSeries {
        if v > expired {
            res += duration
        } else {
            res += v + duration - expired
        }
        expired = v + duration
    }

    return res
}