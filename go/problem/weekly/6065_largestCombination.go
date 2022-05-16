package weekly
import "strconv"

func change(n int) string {
    result := ""
    
    for ; n > 0; n /= 2 {
        lsb := n % 2
        result = strconv.Itoa(lsb) + result
    }
    
    for len(result) < 32 {
        result = "0" + result
    }
    //fmt.Println(result)
    return result
}

func max(x, y int) int {
    if x < y {
        return y
    }
    
    return x
}

func largestCombination(candidates []int) int {
    var ss []string
    for _, v := range candidates {
        ss = append(ss, change(v))
    }
    var res int
    for i := 0; i < 32; i++ {
        tmp := 0
        for j := 0; j < len(ss); j++ {
            if ss[j][i] == '1' {
                tmp++
                // fmt.Println(tmp)
            } 
        }
        res = max(res, tmp)
    }
    return res
}

