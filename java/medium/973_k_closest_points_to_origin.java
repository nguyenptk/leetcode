// https://leetcode.com/problems/kth-largest-element-in-an-array

package java.medium;

import java.util.PriorityQueue;

class KClosestPointsToOrigin {
    public int[][] kClosest(int[][] points, int k) {
        PriorityQueue<int[]> minHeap = new PriorityQueue<>(
                (a, b) -> Integer.compare(
                        (a[0] * a[0] + a[1] * a[1]),
                        (b[0] * b[0] + b[1] * b[1])));

        for (int[] point : points) {
            minHeap.add(point);
        }

        int[][] result = new int[k][2];
        for (int i = 0; i < k; i++) {
            int[] curr = minHeap.poll();
            result[i][0] = curr[0];
            result[i][1] = curr[1];
        }

        return result;
    }
}