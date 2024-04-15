package java.medium;

import java.util.Arrays;
import java.util.Comparator;

class NonOverlappingIntervals {
    public int eraseOverlapIntervals(int[][] intervals) {
        Arrays.sort(intervals, Comparator.comparingInt(a -> a[1]));
        int count = 0;
        int k = Integer.MIN_VALUE;
        for (int[] interval : intervals) {
            int x = interval[0];
            int y = interval[1];

            if (x >= k) {
                k = y;
            } else {
                count++;
            }
        }
        return count;
    }
}
