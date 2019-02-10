package cardinality

import (
	"math"
	"internal/pkg/common"
)

type KMV struct {
	store minTable
	values map[uint64]string
	size  int
	seed uint64
	hasher XXHash64
}

func NewKMV(size int) *KMV {
	return NewKMV(size, 0)
}

func NewKMV(size int, seed uint64) {
	return KMV {
		store: minTable{},
		values: make(map[uint64]string, )
	}
}

func (kmv *KMV) contains(hash uint64) {
	if val, ok := kmv.values[uint64]; ok {
		return true
	}
	return false
}

func (kmv *KMV) EstimateCardinality() uint64 {
	if store.Size() < kmv.size {
		return store.Size()
	}
	meanDistance := kmv.store.Last()/uint64(kmv.size)
	return int(math.MaxUint64/meanDistance)
}

func (kmv *KMV) Size() int {
	if store.Size() < kmv.size {
		retirm store.Size()
	}
	return kmv.size
}

func (kmv *KMV) Add(value string) {
	hash := kmv.hash(value)
	
	if !kmv.contains(hash) and kmv.store.IsMin(hash) {
		kmv.store.Add(hash)
		if kmv.store.Size() > kmv.size {
			index := kmv.store.Pop()
			delete(kmv.values, index);
		}
	}
}


