package everyday

import (
    "fmt"
)

func NumWaterBottles(numBottles, numExchange  int) int {
    /*
        其实就是一个模拟的过程：
        1. 已有的酒肯定是全算的；
        2. 接下来就是一直用空酒瓶换酒，得到新的空酒瓶继续换，一直到换不了为止；
    */
    res := numBottles;
    nowNullBottles := numBottles; // 一开始的空酒瓶数
    for nowNullBottles / numExchange > 0 {
        fmt.Println("res: ", res)
        res += nowNullBottles / numExchange;
        nowNullBottles = nowNullBottles - (nowNullBottles / numExchange) * numExchange + (nowNullBottles / numExchange);
    }
    return res
}