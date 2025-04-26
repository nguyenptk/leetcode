// https://leetcode.com/problems/product-of-the-last-k-numbers/
package medium

type ProductOfNumbers struct {
	arr []int
}

func ConstructorProductOfNumbers() ProductOfNumbers {
	return ProductOfNumbers{
		arr: make([]int, 0),
	}
}

func (this *ProductOfNumbers) Add(num int) {
	this.arr = append(this.arr, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	product := 1
	for i := len(this.arr) - 1; i >= 0; i-- {
		if k > 0 {
			product *= this.arr[i]
			k--
		} else {
			break
		}
	}
	return product
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
