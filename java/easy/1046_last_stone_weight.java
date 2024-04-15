// https://leetcode.com/problems/last-stone-weight

package java.easy;

import java.util.PriorityQueue;

class LastStoneWeight {
    public int lastStoneWeight(int[] stones) {

        PriorityQueue<Integer> heap = new PriorityQueue<>();

        for (int i = 0; i < stones.length; i++) {
            heap.offer(-stones[i]);
        }

        while (heap.size() > 1) {
            int y = heap.poll();
            int x = heap.poll();

            int newStone = -y + x;
            heap.offer(-newStone);
        }

        int result = -heap.poll();

        return result;
    }
}