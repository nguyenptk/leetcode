package java.medium;

class JumpGame {
    public boolean canJump(int[] nums) {
        int n = nums.length;
        int lastPos = n;

        for (int i = n - 1; i >= 0; i--) {
            if (i+ nums[i] >= lastPos) {
                lastPos = i;
            }
        }

        return lastPos == 0;
    }
}
