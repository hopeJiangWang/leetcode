#include <unordered_set>
#include <iostream>

using namespace std;

int lengthOfLongestSubstring(string s) {
    // 异常情况：字符串为空
    if (s.empty() || s.size() == 0) {
        return 0;
    }

    // 双指针滑动查找：左指针为起点，右指针滑动到最远的不重复字符处，记录当前的子串，储存起来
    // 随后，获取下一个子串，与当前子串比较，更新获取最长的子串
    int left = 0, right = 0;
    unordered_set<char> set;
    int res = 0;
    int now_len = 0;
    while (right < s.size()) {
        cout << "left: " << left << "right: " << right << "res: " << res << endl;
        if (!set.count(s[right])) {
            set.insert(s[right]);
            right++;
        } else {
            left++;
            right = left;
            set.clear();    // 注意，这里需要清空，防止对后续的判断影响
        }
        res = max(res, right-left);
    }
    return res;
}

int main() {
    string str;
 
    while (cin >> str)
    {
        cout << lengthOfLongestSubstring(str) << endl;
    }
}