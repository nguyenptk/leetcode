// https://leetcode.com/problems/subarray-sum-equals-k/

package java.medium;

import java.util.HashMap;

class SubArraySumEquals {
    public int subarraySum(int[] nums, int k) {
        HashMap<Integer, Integer> prefixSums = new HashMap<>();
        prefixSums.put(0, 1);

        int sum = 0;
        int result = 0;

        for (int i = 0; i < nums.length; i++) {
            sum += nums[i];
            int diff = sum - k;
            if (prefixSums.containsKey(diff)) {
                result += prefixSums.get(diff);
            }
            prefixSums.put(sum, prefixSums.get(sum) + 1);
        }

        return result;
    }
}