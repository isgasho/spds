package sample

type Sampler interface {
	Sample() []string
	ElementsAdded() uint64
	Size() int
	Add(value string)
}
