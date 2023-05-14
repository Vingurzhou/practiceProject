//
// Created by 周文喆 on 2023/5/14.
//
#include <iostream>
#include "twoSum.cpp"

int main() {
    solution1::Solution s1;
    vector<int> nums = {2, 7, 11, 15};
    vector<int> res = s1.twoSum(nums, 9);
    std::copy(res.begin(), res.end(), std::ostream_iterator<int>(std::cout, " "));
    return 0;
}
