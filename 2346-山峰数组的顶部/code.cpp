// 1.遍历数组
class Solution {
public:
    int peakIndexInMountainArray(vector<int>& arr) {
        int hill, len = arr.size();
        for(int i = 1; i < len - 1; i ++) {
            if(arr[i] > arr[i - 1] && arr[i] > arr[i + 1]) return i;
        }
        return 0;
    }
};

// 2.二分查找
class Solution {
public:
    int peakIndexInMountainArray(vector<int>& arr) {
        int left = 1, right = arr.size() - 1, mid;
        while(left <= right) {
            mid = (left + right) / 2;
            int mid_val = arr[mid], left_val = arr[mid - 1], right_val = arr[mid + 1];
            if(mid_val > left_val && mid_val > right_val) break;
            else if(mid_val < left_val) right = mid - 1;
            else left = mid + 1;
        }
        return mid;
    }
};