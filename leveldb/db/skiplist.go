package db

import (
	"math/rand"
	"sync/atomic"
)

const KMaxHeight = 12

type SkipList struct {
	// 需要一个compartor
	rnd_        *rand.Rand
	max_height_ uint32
	head_       node
}

func New() *SkipList {
	sl := &SkipList{
		// TODO
	}
	return sl
}

func (sl *SkipList) getMaxHeight() uint32 {
	return atomic.LoadUint32(&sl.max_height_)
}

func (sl *SkipList) newNode(key []byte, height int) *node {

}

func (sl *SkipList) Equals(a []byte, b []byte) bool {

}

func (sl *SkipList) randomHeight() (height int32) {
	const kBranching = 4
	height = 1
	for height < KMaxHeight && sl.rnd_.Int31n(kBranching) == 0 {
		height++
	}
}

func (sl *SkipList) keyIsAfterNode(key []byte, n *node) bool {

}

func (sl *SkipList) findGreaterOrEqual(key []byte, prev **node) *node {

}

func (sl *SkipList) findLessThan(key []byte) *node {

}

func (sl *SkipList) findLast() *node {

}

type node struct {
	key   []byte
	next_ []atomic.Value
}

func (nd *node) next(n int) *node {
	return nd.next_[n].Load().(*node)
}

func (nd *node) setNext(n int, x *node) {
	nd.next_[n].Store(x)
}

type Iterator struct {
	list_ *SkipList
	node_ *node
}

func newIterator(list *SkipList) *Iterator {
	iter := &Iterator{
		list,
		nil,
	}
	return iter
}

func (iter *Iterator) Valid() bool {

}

func (iter *Iterator) Key() []byte {

}

func (iter *Iterator) Next() {

}

func (iter *Iterator) Prev() {

}

func (iter *Iterator) Seek(target []byte) {

}

func (iter *Iterator) SeekToFirst() {

}

func (iter *Iterator) SeekToLast() {

}
