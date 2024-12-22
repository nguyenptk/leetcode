package easy

type Logger struct {
	data map[string]int
}

func ConstructorLogger() Logger {
	return Logger{
		data: make(map[string]int),
	}
}

func (f *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if _, ok := f.data[message]; !ok {
		f.data[message] = timestamp
		return true
	}

	old := f.data[message]
	if timestamp-old >= 10 {
		f.data[message] = timestamp
		return true
	}

	return false
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
