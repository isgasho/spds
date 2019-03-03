package sample

import (
	"spds/internal/pkg/common"
)

type MinHash struct {
	store        common.MinTable
	values       map[uint64]string
	size         int
	seed         uint64
	totalCounter uint64
}

func (mh *MinHash) Sample() []string {
	return nil
}

func (mh *MinHash) ElementsAdded() uint64 {
	return 0
}

func (mh *MinHash) Size() int {
	return 0
}

func (mh *MinHash) Add() {

}
