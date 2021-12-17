#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
    int numWaterBottles(int numBottles, int numExchange) {
        /*
            其实就是一个模拟的过程：
            1. 已有的酒肯定是全算的；
            2. 接下来就是一直用空酒瓶换酒，得到新的空酒瓶继续换，一直到换不了为止；
        */
        int res = numBottles;
        int nowNullBottles = numBottles; // 一开始的空酒瓶数
        while(nowNullBottles / numExchange) {
            cout << "res: " << res << endl;
            res += nowNullBottles / numExchange;
            nowNullBottles = nowNullBottles - (nowNullBottles / numExchange) * numExchange + (nowNullBottles / numExchange);
        }
        return res;
    }
};

int main() {
    Solution s;

    cout << s.numWaterBottles(15, 4) << endl;
}