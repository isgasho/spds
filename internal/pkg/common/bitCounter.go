package common

type BitCounter struct {
	counters []uint64
}

func NewBitCounter(size int) *BitCounter {
	return &BitCounter{counters: make([]uint64, size, size)}
}

func (bc *BitCounter) Get(index int) uint64 {
	if index >= len(bc.counters) {
		// Return error
	}
	return bc.counters[index]
}

func (bc *BitCounter) Set(index int, value uint64) {
	if index >= len(bc.counters) {
		// Return error
	}
	bc.counters[index] = value
}
