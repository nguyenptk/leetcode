// https://leetcode.com/problems/snapshot-array/
package medium

type Ele struct {
	val    int
	snapId int
}

type SnapshotArray struct {
	data   [][]Ele
	snapId int
}

func ConstructorSnapshotArray(length int) SnapshotArray {
	data := make([][]Ele, length)
	for i := range data {
		data[i] = []Ele{{0, 0}}
	}
	return SnapshotArray{
		data:   data,
		snapId: 0,
	}
}

func (sa *SnapshotArray) Set(index int, val int) {
	arr := sa.data[index]
	last := len(arr) - 1
	if arr[last].snapId == sa.snapId {
		sa.data[index][last].val = val
	} else {
		sa.data[index] = append(arr, Ele{val, sa.snapId})
	}
}

func (sa *SnapshotArray) Snap() int {
	id := sa.snapId
	sa.snapId++
	return id
}

func (sa *SnapshotArray) Get(index int, snapId int) int {
	arr := sa.data[index]
	l, r := 0, len(arr)-1
	for l <= r {
		m := l + (r-l)/2
		if arr[m].snapId > snapId {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return arr[l-1].val
}

// Time complexity: O(logn)
// Space complexity: O(n+k)

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
