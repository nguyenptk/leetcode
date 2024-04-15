package java.medium;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;

class NumberOfConnectedComponentsInAnUndirectedGraph {
    public int countComponents(int n, int[][] edges) {
        // Build graph from edges
        Map<Integer,List<Integer>> graph = new HashMap<>();
        for (int[] edge : edges) {
            int x = edge[0];
            int y = edge[1];
            if (graph.get(x) == null) graph.put(x, new ArrayList<>());
            if (graph.get(y) == null) graph.put(y, new ArrayList<>());
            graph.get(x).add(y);
            graph.get(y).add(x);
        }

        int result = 0;
        Set<Integer> visited = new HashSet<>();
        for (int i = 0; i < n; i++) {
            if (!visited.contains(i)) {
                explore(graph, i, visited);
                result++;
            }
        }

        return result;
    }

    private void explore(Map<Integer,List<Integer>> graph, int current, Set<Integer> visited) {
        if (visited.contains(current)) return;
        visited.add(current);
        if (graph.get(current) != null) {
            for (int neighbor : graph.get(current)) {
                explore(graph, neighbor, visited);
            }
        }
    }
}
