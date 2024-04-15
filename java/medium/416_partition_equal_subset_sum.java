// https://leetcode.com/problems/partition-equal-subset-sum/

package java.medium;

import java.util.Arrays;

class PartitionEqualSubsetSum {
    public boolean canPartition(int[] nums) {
        int sum = Arrays.stream(nums).sum();
        if (sum % 2 != 0) {
            return false;
        }

        int target = sum / 2;
        boolean[] dp = new boolean[target];
        dp[0] = true;

        for (int n : nums) {
            for (int i = target; i >= n; i--) {
                if (dp[i - n] == true) {
                    if (i == target) {
                        return true;
                    }
                    dp[i] = true;
                }
            }
        }

        return false;
    }
}

