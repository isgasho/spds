package sampling

import "spds/internal/pkg/common"

type Affirmative struct {
	store        common.AdaptiveTable
	values       map[uint64]string
	size         int
	seed         uint64
	totalCounter uint64
}

func (aff *Affirmative) Sample() []string {

}

func (aff *Affirmative) ElementsAdded() uint64 {

}

func (aff *Affirmative) Size() int {

}

func (aff *Affirmative) Add() {

}
