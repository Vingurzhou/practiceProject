#include <iostream>
#include <vector>
#include "Solution.h"

int main() {
    std::vector<int> nums = {2, 7, 11, 15};
    int target = 9;

    Solution solution;
    std::vector<int> result = solution.twoSum(nums, target);

    std::cout  << result[0] << ", " << result[1] << std::endl;

    return 0;
}
