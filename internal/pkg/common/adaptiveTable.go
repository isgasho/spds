package common

type AdaptiveTable struct {
	values      []uint64
	set         map[uint64]bool
	initialSize int
	maxSize     int
	threshold   int
}

func NewAdaptiveTable(size, maxSize int) AdaptiveTable {
	return AdaptiveTable{initialSize: size, maxSize: maxSize, threshold: size}
}

func (at *AdaptiveTable) Size() int {
	return len(at.values)
}

func (at *AdaptiveTable) Last() uint64 {
	if len(at.values) == 0 {
		// panic
	}
	lastIndex := len(at.values) - 1
	return at.values[lastIndex]
}

func (at *AdaptiveTable) Pop() uint64 {
	if len(at.values) == 0 {
		// panic
	}

	last := at.Last()
	delete(at.set, last)

	at.values = at.values[:len(at.values)-2]
	return last
}

func (at *AdaptiveTable) Contains(value uint64) bool {
	if _, ok := at.set[value]; ok {
		return true
	}
	return false
}

func (at *AdaptiveTable) IsMin(value uint64) bool {
	if (len(at.values) < at.initialSize || at.Last() > value) && !at.Contains(value) {
		return true
	}
	return false
}

func (at *AdaptiveTable) Insert(value uint64) (int, uint64) {
	if !at.IsMin(value) {
		return -1, 0
	}

	at.values = append(at.values, value)

	index := len(at.values) - 1
	done := false
	for done != true {
		if index == 0 || at.values[index-1] < at.values[index] {
			done = true
		} else if at.values[index-1] > at.values[index] {
			at.values[index-1], at.values[index] = at.values[index], at.values[index-1]
			index--
		} else {
			done = true
		}
	}

	if index <= at.threshold {
		return index, 0
	}
	return index, at.Pop()
}

func (at *AdaptiveTable) UpdateThreshold(percentage float64) int {
	lenght := float64(len(at.values))
	newThreshold := percentage / 100.0 * lenght
	if newThreshold <= 0 || newThreshold >= lenght {
		// panic
	}
	at.threshold = int(newThreshold)
	return at.threshold
}
