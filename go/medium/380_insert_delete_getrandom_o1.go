// https://leetcode.com/problems/insert-delete-getrandom-o1/
package medium

import "math/rand"

type RandomizedSet struct {
	m map[int]int
	l []int
}

func ConstructorRandomizedSet() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]int),
		l: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.m[val] = len(this.l)
	this.l = append(this.l, val)

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.m[val]
	if !ok {
		return false
	}

	// move the last element
	last := this.l[len(this.l)-1]
	this.l[idx] = last
	this.m[last] = idx

	// delete the last element
	delete(this.m, val)
	this.l = this.l[:len(this.l)-1]

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.l[rand.Intn(len(this.l))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
