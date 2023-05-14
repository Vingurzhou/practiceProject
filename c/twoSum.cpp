//
// Created by 周文喆 on 2023/5/14.
//
#include <vector>
#include <algorithm>

using namespace std;
namespace solution1 {
    class Solution {
    public:
        vector<int> twoSum(vector<int>& nums, int target) {
            vector<int> sorted_nums = nums;
            sort(sorted_nums.begin(), sorted_nums.end());
            int left = 0, right = sorted_nums.size() - 1;
            while (left < right) {
                int sum = sorted_nums[left] + sorted_nums[right];
                if (sum == target) {
                    break;
                } else if (sum < target) {
                    ++left;
                } else {
                    --right;
                }
            }
            vector<int> res;
            for (int i = 0; i < nums.size(); ++i) {
                if (nums[i] == sorted_nums[left] || nums[i] == sorted_nums[right]) {
                    res.push_back(i);
                }
            }
            return res;
        }
    };
}
