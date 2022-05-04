package everyday

func findTheWinner(n int, k int) int {
    /*
        约瑟夫环问题
    */
    if n <= 1{
        return n
    } 
    ans := (findTheWinner(n - 1, k) + k) % n

    if ans == 0 {
        ans = n
    }
    return ans


    /*
    q := make([]int, n)
    for i := range q {
        q[i] = i + 1
    }
    for len(q) > 1 {
        for i := 1; i < k; i++ {
            q = append(q, q[0])[1:]
        }
        q = q[1:]
    }
    return q[0]
    */
}