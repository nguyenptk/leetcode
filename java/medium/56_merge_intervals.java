package java.medium;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class MergeIntervals {
    public int[][] merge(int[][] intervals) {
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        List<int[]> result = new ArrayList<>();
        for (int [] interval : intervals) {
            if (result.isEmpty() || result.get(result.size()-1)[1] < interval[0]) {
                result.add(interval);
            }
            else {
                result.get(result.size()-1)[1] = Math.max(result.get(result.size()-1)[1], interval[1]);
            }
        }
        return result.toArray(new int[result.size()][]);
    }
}
