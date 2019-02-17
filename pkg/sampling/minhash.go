package sampling

type MinHash struct {
	store        common.minTable
	values       map[uint64]string
	size         int
	seed         uint64
	totalCounter uint64
}

func (mh *MinHash) Sample() []string {

}

func (mh *MinHash) ElementsAdded() uint64 {

}

func (mh *MinHash) Size() int {

}

func (mh *MinHash) Add() {

}
