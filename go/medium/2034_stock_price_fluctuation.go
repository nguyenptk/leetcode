// https://leetcode.com/problems/stock-price-fluctuation

package medium

import "container/heap"

type stockQuote struct {
	ts    int
	price int
}

type StockQuoteHeap []stockQuote

func (h StockQuoteHeap) Len() int           { return len(h) }
func (h StockQuoteHeap) Less(i, j int) bool { return h[i].price > h[j].price }
func (h StockQuoteHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *StockQuoteHeap) Push(x interface{}) {
	*h = append(*h, x.(stockQuote))
}
func (h *StockQuoteHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h StockQuoteHeap) Top() stockQuote {
	return h[0]
}

type StockPrice struct {
	last     stockQuote
	max      *StockQuoteHeap
	min      *StockQuoteHeap
	storedTS map[int]int
}

func ConstructorStockPrice() StockPrice {
	max := StockQuoteHeap(make([]stockQuote, 0))
	min := StockQuoteHeap(make([]stockQuote, 0))
	stored := make(map[int]int)
	return StockPrice{max: &max, min: &min, storedTS: stored}
}

func (p *StockPrice) Update(timestamp int, price int) {
	if timestamp >= p.last.ts {
		p.last = stockQuote{ts: timestamp, price: price}
	}
	heap.Push(p.max, stockQuote{ts: timestamp, price: price})
	heap.Push(p.min, stockQuote{ts: timestamp, price: -price})
	p.storedTS[timestamp] = price
}

func (p *StockPrice) Current() int {
	return p.last.price
}

func (p *StockPrice) Maximum() int {
	top := p.max.Top()
	price := p.storedTS[top.ts]
	if price > 0 && price != top.price {
		heap.Pop(p.max)
		return p.Maximum()
	}
	return top.price
}

func (p *StockPrice) Minimum() int {
	top := p.min.Top()
	price := p.storedTS[top.ts]
	if price > 0 && price != -top.price {
		heap.Pop(p.min)
		return p.Minimum()
	}
	return -top.price
}
