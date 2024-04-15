package java.medium;

import java.util.Arrays;

// Brute Force
class Solution {
    private int result = 0;

    public int findTargetSumWays(int[] nums, int S) {
        calculate(nums, S, 0, 0);
        return result;
    }

    private void calculate(int[] nums, int S, int idx, int sum) {
        if (idx == nums.length) {
            if (sum == S) {
                result++;
            }
        } else {
            calculate(nums, S, idx + 1, sum + nums[idx]);
            calculate(nums, S, idx + 1, sum - nums[idx]);
        }
    }
}

// DP 1D
// class TargetSum {
//     public int findTargetSumWays(int[] nums, int S) {
//         int sum = Arrays.stream(nums).sum();
//         if (sum < S || (sum-S) % 2 != 0) {
//             return 0;
//         }

//         int target = (sum - S) / 2;
//         int[] dp = new int[target+1];
//         dp[0] = 1;

//         for (int num : nums) {
//             for (int i = target; i >= 0; i--) {
//                 if (i >= num) {
//                     dp[i] += dp[i-num];
//                 }
//             }
//         }

//         return dp[target];
//     }
// }
