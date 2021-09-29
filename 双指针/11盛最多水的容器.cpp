#include <bits/stdc++.h>

using namespace std;

int maxArea(vector<int>& height) {
    // 数组为空，直接返回0即可
    if (height.empty()) {
        return 0;
    }

    /*
        朴素算法，超时

    // 双指针一直往后查找，并记录当前最大的面积
    int res = 0;
    for (int i = 0; i < height.size(); i++) {
        for (int j = i + 1; j < height.size(); j++) {
            int w = j - i;
            int h = min(height[i], height[j]);
            res = max(res, w * h);
        }
    }
    return res;
    */
    // 左右指针分别指向最左端和最右端，依次选择往中间移动，但是往中间移动会导致宽度减小，这时候就需要
    // 保证移动之后高度增大：即只能选择高度较小的指针往中间移动，才可能能够获取到更大的面积
    int res = 0;
    int left = 0, right = height.size() - 1;
    while (left < right) {
        res = max(res, (right - left) * min(height[right], height[left]));
        if (height[left] <= height[right]) {
            left++;
        } else {
            right--;
        }
    }
    return res;
}

int main() {
    vector<int> v = {4, 3, 2, 1, 4};    // 16
    // vector<int> v = {1,8,6,2,5,4,8,3,7}; // 49

    cout << maxArea(v) << endl;
}