// https://leetcode.com/problems/find-building-where-alice-and-bob-can-meet/
package hard

import "github.com/emirpasic/gods/v2/queues/priorityqueue"

type Query struct {
	Height int
	Index  int
}

func LeftmostBuildingQueries(Heights []int, queries [][]int) []int {
	storeQueries := make([][]Query, len(Heights))
	for i := range storeQueries {
		storeQueries[i] = make([]Query, 0)
	}

	result := make([]int, len(queries))
	for i := range result {
		result[i] = -1
	}

	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		return a.(Query).Height - b.(Query).Height
	})

	for currQuery, query := range queries {
		a, b := query[0], query[1]
		if a < b && Heights[a] < Heights[b] {
			result[currQuery] = b
		} else if a > b && Heights[a] > Heights[b] {
			result[currQuery] = a
		} else if a == b {
			result[currQuery] = a
		} else {
			maxHeight := max(Heights[a], Heights[b])
			maxQuery := max(a, b)
			storeQueries[maxQuery] = append(storeQueries[maxQuery], Query{Height: maxHeight, Index: currQuery})
		}
	}

	for i := 0; i < len(Heights); i++ {
		for pq.Size() > 0 {
			top, _ := pq.Peek()
			if top.(Query).Height >= Heights[i] {
				break
			}
			resolved, _ := pq.Dequeue()
			result[resolved.(Query).Index] = i
		}

		for _, element := range storeQueries[i] {
			pq.Enqueue(element)
		}
	}

	return result
}
