package weekly

import (
	"fmt"
	"sort"
)

func maxConsecutive(bottom int, top int, special []int) int {
    var res int
    sort.Ints(special)
    
    for _, v := range special {
        fmt.Println(v, bottom)
        res = max(res, v-bottom)
        // fmt.Println(tmp, i)
        bottom = v+1
    }
    fmt.Println(top, bottom)
    res = max(res, top-bottom+1)
    
    return res
}
