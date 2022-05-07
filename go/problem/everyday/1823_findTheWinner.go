package everyday

func findTheWinner(n int, k int) int {
    /*
        约瑟夫环问题
        1、每次往同一方向，以固定步长k进行消数。
        2、由于下一次操作的发七点为消除位置的下一个点（即前后两次操作发起点在原序列下标中相差 k）
        3、这时候问题规模会从n变成n-1，原问题等价于findTheWinner(n-1, k) + k
        ==> f(n, m) = (f(n01, m) + m) % n
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