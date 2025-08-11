// https://leetcode.com/problems/design-a-text-editor/
package hard

type TextEditor struct {
	left  []byte
	right []byte
}

func ConstructorTextEditor() TextEditor {
	return TextEditor{
		left:  make([]byte, 0),
		right: make([]byte, 0),
	}
}

func (f *TextEditor) AddText(text string) {
	for i := 0; i < len(text); i++ {
		f.left = append(f.left, text[i])
	}
}

func (f *TextEditor) DeleteText(k int) int {
	count := 0
	for len(f.left) > 0 && k > 0 {
		f.left = f.left[:len(f.left)-1]
		count++
		k--
	}
	return count
}

func (f *TextEditor) CursorLeft(k int) string {
	for len(f.left) > 0 && k > 0 {
		c := f.left[len(f.left)-1]
		f.left = f.left[:len(f.left)-1]
		f.right = append(f.right, c)
		k--
	}
	// returning the last min(10, len) characters to the left of the cursor
	return f.cursorShiftString()
}

func (f *TextEditor) CursorRight(k int) string {
	for len(f.right) > 0 && k > 0 {
		c := f.right[len(f.right)-1]
		f.right = f.right[:len(f.right)-1]
		f.left = append(f.left, c)
		k--
	}
	// returning the last min(10, len) characters to the left of the cursor
	return f.cursorShiftString()
}

func (f *TextEditor) cursorShiftString() string {
	total := len(f.left)
	start := total - 10
	if start < 0 {
		start = 0
	}
	return string(f.left[start:])
}

/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */
