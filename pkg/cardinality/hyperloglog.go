package cardinality

import (
	"math"
	"spds/internal/pkg/common"

	"github.com/cespare/xxhash"
)

type HLL struct {
	store        common.BitCounter
	size         uint64
	indexSize    uint64
	alpha        float64
	seed         uint64
	totalCounter uint64
}

func NewHLL(size int) *HLL {
	indexSize := math.Log2(float64(size))

	// TODO: Check that size is a power of 2
	if indexSize != float64(int(indexSize)) {

	}

	alpha := (0.7213 / (1.0 + 1.079/float64(size)))
	return &HLL{
		store:        common.NewBitCounter(size),
		size:         uint64(size),
		indexSize:    uint64(indexSize),
		alpha:        alpha,
		seed:         0,
		totalCounter: 0,
	}
}

func (hll *HLL) hash(s string) uint64 {
	return xxhash.Sum64String(s)
}

func (hll *HLL) getIndex(bits uint64) int {
	return int(bits >> (64 - hll.indexSize))
}

func (hll *HLL) getDepth(bits uint64) uint64 {
	offset := bits & (math.MaxUint64 >> hll.indexSize)
	return 64 - hll.indexSize - uint64(math.Log2(float64(offset)))
}

func (hll *HLL) EstimateCardinality() uint64 {
	floatSize := float64(hll.size)
	numerator := hll.alpha * math.Pow(floatSize, 2)
	denominator := float64(0)

	for i := 0; i < hll.store.Size(); i++ {
		denominator += (1.0 / math.Pow(2, float64(hll.store.Get(i))))
	}

	estimation := numerator / denominator

	// Correcting bias
	if estimation <= 5/2*floatSize {
		numberOfZeros := float64(hll.store.ZeroCounter())
		estimation = floatSize * math.Log(floatSize/numberOfZeros)
	} else if estimation > 1.0/30*math.Log(1-estimation/math.Pow(2, 32)) {
		estimation = -math.Pow(2, 32) * math.Log2(1-estimation/math.Pow(2, 32))
	}

	return uint64((1.0 / hll.alpha) * estimation)
}

func (hll *HLL) ElementsAdded() uint64 {
	return hll.totalCounter
}

func (hll *HLL) Size() int {
	return hll.store.Size()
}

func (hll *HLL) Add(value string) {
	hll.totalCounter++
	hash := hll.hash(value)

	index := hll.getIndex(hash)
	depth := hll.getDepth(hash)

	if hll.store.Get(index) < depth {
		hll.store.Set(index, depth)
	}
}
