#include <algorithm>

/*
 * 思路：排序后从后往前遍历
 * 写比较函数的时候，需要用'const vector<int> &'，否则超时
 */
class Solution {
public:
    int numberOfWeakCharacters(vector<vector<int>>& properties) {
        sort(properties.begin(), properties.end(), [](const vector<int> & n1, const vector<int> &n2) {
            return n1[0] != n2[0] ? (n1[0] < n2[0]) : (n1[1] > n2[1]);
        });
        int max_def = properties[properties.size() - 1][1], max_atta = properties[properties.size() - 1][0];
        int cur = properties.size() - 2, ans = 0;
        while(cur >= 0 && properties[cur][0] == max_atta) {
            max_def = properties[cur][1] > max_def ? properties[cur][1] : max_def;
            cur --;
        }

        while(cur >= 0) {
            if(properties[cur][1] < max_def) ans ++;
            else if(properties[cur][1] > max_def) {
                max_def = properties[cur][1];
            }
            cur --;
        }
        return ans;
    }
};