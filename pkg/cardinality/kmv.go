package cardinality

import (
	"hash"
	"math"
	"xxhash"

	"spds/internal/pkg/common"
)

type KMV struct {
	store        common.MinTable
	values       map[uint64]string
	size         int
	seed         uint64
	hasher       hash.Hash64
	totalCounter uint64
}

func NewKMV(size int) KMV {
	return NewKMVWithSeed(size, 0)
}

func NewKMVWithSeed(size int, seed uint64) KMV {
	hasher := xxhash.New()
	return KMV{
		store:  common.MinTable{},
		values: make(map[uint64]string, size),
		size:   size,
		seed:   seed,
		hasher: hasher}
}

func (kmv *KMV) contains(hash uint64) bool {
	if _, ok := kmv.values[hash]; ok {
		return true
	}
	return false
}

func (kmv *KMV) hash(s string) uint64 {
	return xxhash.Sum64String(s)
}

func (kmv *KMV) EstimateCardinality() uint64 {
	if kmv.store.Size() < kmv.size {
		return uint64(kmv.store.Size())
	}
	meanDistance := kmv.store.Last() / uint64(kmv.size)
	return math.MaxUint64 / meanDistance
}

func (kmv *KMV) ElementsAdded() uint64 {
	return kmv.totalCounter
}

func (kmv *KMV) Size() int {
	if kmv.store.Size() < kmv.size {
		return kmv.store.Size()
	}
	return kmv.size
}

func (kmv *KMV) Add(value string) {
	kmv.totalCounter++
	hash := kmv.hash(value)

	if !kmv.contains(hash) && kmv.store.IsMin(hash) {
		kmv.store.Insert(hash)
		if kmv.store.Size() > kmv.size {
			index := kmv.store.Pop()
			delete(kmv.values, index)
		}
	}
}
