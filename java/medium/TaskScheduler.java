package java.medium;

import java.util.Collections;
import java.util.LinkedList;
import java.util.PriorityQueue;
import java.util.Queue;

public class TaskScheduler {
    class Pair<K, V> {
        K key;
        V value;

        public Pair(K key, V value) {
            this.key = key;
            this.value = value;
        }

        public K getKey() {
            return key;
        }
    }

    public int leastInterval(char[] tasks, int n) {

        // Build frequency map
        int[] freq = new int[26];
        for (char c : tasks) {
            freq[c-'a']++;
        }

        PriorityQueue<Integer> pq = new PriorityQueue<>(Collections.reverseOrder());
        for (int i = 0; i < 26; i++) {
            if (freq[i] > 0) {
                pq.offer(freq[i]);
            }
        }

        int time = 0;
        Queue<Pair<Integer, Integer>> q = new LinkedList<>();
        while (!pq.isEmpty() || !q.isEmpty()) {
            time++;
            if (!pq.isEmpty()) {
                int val = pq.poll();
                val--;
                if (val > 0) {
                    q.offer(new Pair<>(val, time+n));
                }
            }
            if (!q.isEmpty() && q.peek().value <= time) {
                pq.add(q.poll().getKey());
            }
        }

        return time;
    }
}
