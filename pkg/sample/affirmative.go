package sample

import (
	"math"
	"spds/internal/pkg/common"

	"github.com/cespare/xxhash"
)

type Affirmative struct {
	store        common.AdaptiveTable
	values       map[uint64]string
	size         int
	seed         uint64
	totalCounter uint64
}

func NewAffirmative(size int) *Affirmative {
	return &Affirmative{
		store:        common.NewAdaptiveTable(size, math.MaxInt32),
		values:       make(map[uint64]string, size),
		size:         size,
		seed:         0,
		totalCounter: 0,
	}
}
func (aff *Affirmative) contains(hash uint64) bool {
	if _, ok := aff.values[hash]; ok {
		return true
	}
	return false
}

func (aff *Affirmative) hash(s string) uint64 {
	return xxhash.Sum64String(s)
}

func (aff *Affirmative) Sample() []string {
	var sample []string
	for _, value := range aff.values {
		sample = append(sample, value)
	}

	return sample
}

func (aff *Affirmative) ElementsAdded() uint64 {
	return aff.totalCounter
}

func (aff *Affirmative) Size() int {
	return len(aff.values)
}

func (aff *Affirmative) Add(value string) {
	aff.totalCounter++
	hash := aff.hash(value)

	if !aff.contains(hash) && aff.store.IsMin(hash) {
		aff.values[hash] = value
		_, oldHash := aff.store.Insert(hash)
		if oldHash != 0 {
			delete(aff.values, oldHash)
		}
	}
}
