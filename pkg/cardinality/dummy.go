package cardinality

type Dummy struct {
	values map[string]bool
}

func (d *Dummy) EstimateCardinality() uint64 {
	return len(d.values)
}

func (d *Dummy) Size() int {
	return len(d.values)
}

func (d *Dummy) Add(value string) {
	d.values[value] = true
}
