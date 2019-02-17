package common

type MinTable struct {
	values      []uint64
	initialSize int
}

func NewMinTable(size int) MinTable {
	return MinTable{initialSize: size}
}

func (mt *MinTable) Size() int {
	return len(mt.values)
}

func (mt *MinTable) Last() uint64 {
	if len(mt.values) == 0 {
		// return error
	}
	lastIndex := len(mt.values) - 1
	return mt.values[lastIndex]
}

func (mt *MinTable) Pop() uint64 {
	if len(mt.values) == 0 {
		// return error
	}
	last := mt.Last()
	mt.values = mt.values[:len(mt.values)-2]
	return last
}

func (mt *MinTable) Contains(value uint64) bool {
	if len(mt.values) == 0 || mt.Last() < value {
		return false
	}

	for _, element := range mt.values {
		if element == value {
			return true
		}
	}

	return false
}

func (mt *MinTable) IsMin(value uint64) bool {
	if (len(mt.values) < mt.initialSize || mt.Last() > value) && !mt.Contains(value) {
		return true
	}
	return false
}

func (mt *MinTable) Insert(value uint64) int {
	if !mt.IsMin(value) {
		return -1
	}

	mt.values = append(mt.values, value)

	index := len(mt.values) - 1
	done := false
	for done != true {
		if index == 0 || mt.values[index-1] < mt.values[index] {
			done = true
		} else if mt.values[index-1] > mt.values[index] {
			mt.values[index-1], mt.values[index] = mt.values[index], mt.values[index-1]
			index--
		} else {
			done = true
		}
	}
	return index
}
