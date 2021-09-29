#include <bits/stdc++.h>

using namespace std;

int longestSubarray(vector<int>& nums, int limit) {
    // 数组为空，直接返回0即可
    if (nums.empty()) {
        return 0;
    }

    // 选定右端点，找到其对应的最左边的端点，满足绝对差值不超过限制
    // 使用有序集合：multiset(nlogn: 往有序集合中增删数据为logn)
    int res = 0;
    int left = 0, right = 0, len = nums.size();
    multiset<int> s;
    while (right < len) {
        s.insert(nums[right]);
        if (*s.rbegin() - *s.begin() > limit) {
            s.erase(s.find(nums[left++]));
        }
        res = max(res, right - left + 1);
        right++;
    }

    return res;
}

 int longestSubarray2(vector<int>& nums, int limit) {
    // 数组为空，直接返回0即可
    if (nums.empty()) {
        return 0;
    }

    // 选定右端点，找到其对应的最左边的端点，满足绝对差值不超过限制
    // 因为法一中查找有序集合耗时logn，可优化为n：
    // 即使用两个单调队列，分别维护区间最大值雨区间最小值：这样的话，就可以直接O(1)获取区间极值
    int ret = 0;
    int left = 0, right = 0, len = nums.size();
    deque<int> queMin;
    deque<int> queMax;
    while (right < len) {
            // 获取当前区间的最大值、最小值
        while (!queMin.empty() && queMin.back() > nums[right]) {
            queMin.pop_back();
        }
        while (!queMax.empty() && queMax.back() < nums[right]) {
            queMax.pop_back();
        }

        // 判断当前区间是否符合标准
        queMax.push_back(nums[right]);
        queMin.push_back(nums[right]);
        while (!queMax.empty() && !queMin.empty() && queMax.front() - queMin.front() > limit) {
            // 如果不符合，则需要将左指针右移；同时，需要将对应的极值队列调整
            if (nums[left] == queMin.front()) {
                queMin.pop_front();
            }
            if (nums[left] == queMax.front()) {
                queMax.pop_front();
            }
            left++;
        }
        ret = max(ret, right - left + 1);
        right++;
    }
    return ret;
}

int main() {
    // vector<int> v = {8,2,4,7};    // 4->2
    vector<int> v = {4,2,2,2,4,4,2,2}; // 49

    cout << longestSubarray2(v, 0) << endl;
    cout << longestSubarray(v, 0) << endl;
} 