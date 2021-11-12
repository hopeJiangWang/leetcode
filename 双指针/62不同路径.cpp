#include <iostream>
#include <vector>

using namespace std;

int uniquePaths(int m, int n) 
{
    /*
        以f[i][j]来表示当前位置需要的步数，则f[n-1][m-1]是我们最终需要获得的结果，f[0][0]=1;
        需要根据当前位置来决定处理计算方式，因为方向都是往下或者往右，所以有以下几种情况：
        1. 如果是只能往右移动：f[i][j] = f[i][j-1] + 1;
        2. 如果只能向下移动：f[i][j] = f[i-1][j] + 1;
        则状态转移方程为：f[i][j] = f[i-1][j] + f[i][j-1];
        初始化条件为f[0][0]=1；
        对于边界，只能沿着边界走：f[0][j] = f[i][0] = 1
    */

    vector<vector<int> > f(m, vector<int>(n, 0));
    for (int i = 0; i < m; ++i) {
        f[i][0] = 1;
    }
    for (int j = 0; j < n; ++j) {
        f[0][j] = 1;
    }
    for (int i = 1; i < m; i++) {
        for (int j = 1; j < n; j++) {
            f[i][j] = f[i - 1][j] + f[i][j - 1];
        }
    }
    return f[m - 1][n - 1];
}

int main() {
    int m = 3, n = 2;
    cout << uniquePaths(m, n) << endl;
}