package top100

func NumTrees(n int) int {
    /*
        假设n个节点存在二叉排序树的个数是f(n)，那么根节点的情况为：1, 2, ... , n
        1为根节点，左子树节点个数为0，右子树为n-1，总的情况会有：f(0)*f(n-1)
        2为根节点，左子树节点个数为1，右子树节点为nn-2，总的情况会有：f(1)*f(n-2)
        ...
        f(n) = f(0)*f(n-1)+f(1)*f(n-2)+...+f(n-1)*f(0)  
    */

    f := make([]int, n+1)
    f[0] = 1
    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            f[i] += f[j-1]*f[i-j]
            // fmt.Println("now: ", f[i])
        }  
    }
    return f[n]
}