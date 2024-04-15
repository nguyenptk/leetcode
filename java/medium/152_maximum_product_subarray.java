// https://leetcode.com/problems/maximum-product-subarray/

package java.medium;

class MaximumProductSubarray {
    public int maxProduct(int[] nums) {
        if (nums.length == 1) {
            return nums[0];
        }

        int result = nums[0];
        int min = 1;
        int max = 1;

        for (int n : nums) {
            int tmp = max * n;
            max = Math.max(n, Math.max(min * n, tmp));
            min = Math.min(n, Math.min(min * n, tmp));

            result = Math.max(result, max);
        }

        return result;
    }
}
