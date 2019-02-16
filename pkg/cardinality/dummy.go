package cardinality

type Dummy struct {
	values       map[string]bool
	totalCounter uint64
}

func NewDummy() Dummy {
	return Dummy{
		values: make(map[string]bool, 1024)}
}

func (d *Dummy) EstimateCardinality() uint64 {
	return uint64(len(d.values))
}

func (d *Dummy) ElementsAdded() uint64 {
	return d.totalCounter
}

func (d *Dummy) Size() int {
	return len(d.values)
}

func (d *Dummy) Add(value string) {
	d.totalCounter++
	d.values[value] = true
}
