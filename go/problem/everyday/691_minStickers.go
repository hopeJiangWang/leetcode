package everyday

/*
输入： stickers = ["with","example","science"], target = "thehat"
输出：3
解释：
我们可以使用 2 个 "with" 贴纸，和 1 个 "example" 贴纸。
把贴纸上的字母剪下来并重新排列后，就可以形成目标 “thehat“ 了。
此外，这是形成目标字符串所需的最小贴纸数量。

*/

func minStickers(stickers []string, target string) int {
    // 标记target的不同队列 2^m，避免重复计算
    m := len(target)
    f := make([]int, 1<<m)
    for i := range f {
        f[i] = -1
    }
    f[0] = 0

    // dp表示不同状态的最小贴纸数：输入mask为某个子序列，输出是拼出该子序列的最小贴纸数
    var dp func(int) int
    dp = func(mask int) int {
        // 如果之前有计算过，那么直接返回即可
        if f[mask] != -1 {
            return f[mask]
        }
        // 初始化为最差的情况，即每个字母都需要从一个贴纸获取
        f[mask] = m + 1
        for _, sticker := range stickers {
            // 需要计算当前mask与sticker字符的最大交集
            // 交集部分有当前sticker贡献，其他部分的最小贴纸数由动归继续计算
            // left标识未被sticker覆盖的字符数目（初始化就是mask）
            left := mask
            cnt := [26]int{}
            // 统计当前贴纸的个字符数目
            for _, c := range sticker {
                cnt[c-'a']++
            }
            for i, c := range target {
                // target当前字符还没有，并且当前贴纸有这个字符
                // 贴纸上的这个字符数--
                // target中这个字符需要标记已经有了
                if mask>>i&1 == 1 && cnt[c-'a'] > 0 {
                    cnt[c-'a']--
                    left ^= 1 << i
                }
            }
            // 遍历完之后，需要取所有贴纸数的最小值作为本次规划的结果
            if left < mask {
                f[mask] = min(f[mask], dp(left)+1)
            }
        }
        return f[mask]
    }
    // 需要凑成target，最终有 1<<m - 1(1: target[i]已经凑成) 
    ans := dp(1<<m - 1)
    // 合法，则认为能成功
    if ans <= m {
        return ans
    }
    return -1
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
