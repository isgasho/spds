package cardinality

type CardinalityEstimator interface {
	EstimateCardinality() uint64
	Size() int
	Add(value string)
}
