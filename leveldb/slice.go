package leveldb

import (
	"bytes"
)

type Slice struct {
	data_ []byte
}

func New() *Slice {
	return &Slice{
		[]byte(""),
	}
}
func NewByLen(n int) *Slice {
	return &Slice{
		make([]byte, n),
	}
}

func NewByOther(key *Slice) *Slice {
	res := NewByLen(len(key.data_))
	copy(res.data_, key.data_)
	return res
}

func NewByDataAndLen(data_ []byte, n int) *Slice {
	res := Slice{
		make([]byte, n),
	}
	copy(res.data_, data_)
	return &res
}
func NewByData(data_ []byte) *Slice {
	return NewByDataAndLen(data_, len(data_))
}
func NewByString(s string) *Slice {
	return &Slice{
		[]byte(s),
	}
}
func (slice *Slice) Empty() bool {
	return len(slice.data_) == 0
}

func (slice *Slice) Size() int {
	return len(slice.data_)
}

func (slice *Slice) Data() []byte {
	return slice.data_
}

func (slice *Slice) Get(n int) byte {
	return slice.data_[n]
}
func (slice *Slice) Clear() {
	slice.data_ = slice.data_[0:0]
}
func (slice *Slice) RemovePrefix(n int) {
	// 判断n<= size
	slice.data_ = slice.data_[n:len(slice.data_)]
}
func (slice *Slice) ToString() string {
	return string(slice.data_[:])
}

// Three-way comparison.  Returns value:
  //   <  0 iff "*this" <  "b",
  //   == 0 iff "*this" == "b",
  //   >  0 iff "*this" >  "b"
func (slice *Slice) Compare(b *Slice) int {
	var min_len int 
	if len(slice.data_) < len(b.data_) {
		min_len = len(slice.data_)
	} else {
		min_len = len(b.data_)
	}
	r := bytes.Compare(slice.data_[:min_len], b.data_[:min_len])
	if r == 0 {
		if len(slice.data_) < len(b.data_) {
			r = -1
		} else if len(slice.data_) > len(b.data_) {
			r = +1
		}
	}
	return r
}

func (slice *Slice) StartsWith(b *Slice) bool {
	min_len := len(b.data_)
	return (len(slice.data_) >= len(b.data_)) && (bytes.Compare(slice.data_[:min_len], b.data_[:min_len]) == 0)
}

func Equals(a *Slice, b *Slice) bool {
	return ((len(a.data_) == len(b.data_)) && (bytes.Compare(a.data_, b.data_) == 0))
}