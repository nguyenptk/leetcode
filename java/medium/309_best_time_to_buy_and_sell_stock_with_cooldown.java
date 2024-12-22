// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

package java.medium;

import java.util.HashMap;
import java.util.Map;

class BestTimeToBuyAndSellStockWithCooldown {
    public int maxProfit(int[] prices) {
        Map<String,Integer> cache = new HashMap<>();
        return dfs(prices, cache, 0, true);
    }

    private int dfs(
        int[] prices,
        Map<String,Integer> cache,
        int index,
        boolean buying
    ) {
        if (index >= prices.length) {
            return 0;
        }
        String key = index + "-" + buying;

        if (cache.containsKey(key)) {
            return cache.get(key);
        }

        int cooldown = dfs(prices, cache, index+1, buying);
        int buyOrSell = Integer.MIN_VALUE;

        if (buying) {
            buyOrSell = dfs(prices, cache, index+1, !buying) - prices[index];
        } else {
            buyOrSell = dfs(prices, cache, index+2, !buying) + prices[index];
        }

        cache.put(key, Math.max(buyOrSell, cooldown));

        return cache.get(key);
    }
}