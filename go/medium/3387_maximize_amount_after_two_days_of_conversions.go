// https://leetcode.com/problems/maximize-amount-after-two-days-of-conversions/
package medium

type pairCurrency struct {
	to   string
	rate float64
}

func MaxAmount(initialCurrency string, pairs1 [][]string, rates1 []float64, pairs2 [][]string, rates2 []float64) float64 {
	var buildGraph func(pairs [][]string, rates []float64) map[string][]pairCurrency

	buildGraph = func(pairs [][]string, rates []float64) map[string][]pairCurrency {
		graph := make(map[string][]pairCurrency)
		for i, p := range pairs {
			from, to := p[0], p[1]
			rate := rates[i]
			graph[from] = append(graph[from], pairCurrency{to, rate})
			graph[to] = append(graph[to], pairCurrency{from, 1.0 / rate})
		}
		return graph
	}

	day1Graph := buildGraph(pairs1, rates1)
	day2Graph := buildGraph(pairs2, rates2)

	day1Amounts := bfsMax(day1Graph, initialCurrency)

	maxAmount := 1.0
	for currency, amt := range day1Amounts {
		day2Amounts := bfsMax(day2Graph, currency)
		if val, ok := day2Amounts[initialCurrency]; ok {
			maxAmount = max(maxAmount, amt*val)
		}
	}
	return maxAmount
}

func bfsMax(graph map[string][]pairCurrency, start string) map[string]float64 {
	q := []string{start}
	amounts := map[string]float64{start: 1.0}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		currAmt := amounts[curr]

		for _, edge := range graph[curr] {
			nextAmt := currAmt * edge.rate
			if nextAmt > amounts[edge.to] {
				amounts[edge.to] = nextAmt
				q = append(q, edge.to)
			}
		}
	}
	return amounts
}
