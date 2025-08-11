// https://leetcode.com/problems/design-spreadsheet/
package medium

import (
	"strconv"
	"strings"
	"unicode"
)

type Spreadsheet struct {
	m map[string]int // cell to value
}

func ConstructorSpreadsheet(rows int) Spreadsheet {
	return Spreadsheet{m: make(map[string]int)}
}

func (f *Spreadsheet) SetCell(cell string, value int) {
	f.m[cell] = value
}

func (f *Spreadsheet) ResetCell(cell string) {
	delete(f.m, cell)
}

func (f *Spreadsheet) GetValue(formula string) int {
	idxOp := strings.Index(formula, "+")
	val1 := f.extractValue(formula[1:idxOp])
	val2 := f.extractValue(formula[idxOp+1:])
	return val1 + val2
}

func (f *Spreadsheet) extractValue(operand string) int {
	if unicode.IsLetter(rune(operand[0])) {
		return f.m[operand]
	}
	val, _ := strconv.Atoi(operand)
	return val
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */
