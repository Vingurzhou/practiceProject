//
// Created by 周文喆 on 2023/5/14.
//

#ifndef TWOSUM_SOLUTION_H
#define TWOSUM_SOLUTION_H


class Solution {
public:
    vector<int> twoSum(vector<int> &nums, int target) {
        int n = nums.size();
        for (int i = 0; i < n; ++i) {
            for (int j = i + 1; j < n; ++j) {
                if (nums[i] + nums[j] == target) {
                    return {i, j};
                }
            }
        }
        return {};
    }
};


#endif //TWOSUM_SOLUTION_H
