package db 

import (
	"github.com/Lchangliang/goleveldb/leveldb/slice"
	"github.com/Lchangliang/goleveldb/leveldb/util/coding"
)
type ValueType int 

const (
	kTypeDeletion ValueType = iota
	kTypeValue
)

var kValueTypeForSeek = kTypeValue

type SequenceNumber uint64

var kMaxSequenceNumber = ((uint64(0x1) << 56) -1)

type ParsedInternalKey struct {
	userKey *Slice
	sequence SequenceNumber
	valueType ValueType
}

func NewParsedInternalKey(userKey *Slice, seq SequenceNumber, t ValueType) *ParsedInternalKey {
	return &ParsedInternalKey {
		userKey,
		seq,
		t
	}
}

func InternalKeyEncodingLength(key *ParsedInternalKey) int {
	return key.userKey.Size() + 8
}

type InternalKey struct {
	req_ string
}
func PacksequenceAndType(seq uint64, t ValueType) uint64 {
	return (seq << 8) | t
}
func AppendInternalKey(key *ParsedInternalKey) string {
	var result string
	result += string(key.userKey.Data())
	result = coding.PutFixed64(result, PacksequenceAndType(key.sequence, key.valueType))
	return result
}
func NewInternalKey(user_key *Slice, s SequenceNumber, t ValueType) *InternalKey {
	return &InternalKey {
		AppendInternalKey(NewParsedInternalKey(user_key, s, t))
	}
}	
func (ikey *InternalKey) DecoFrom(s *Slice) bool {
	req_ = s.ToString()
	return len(ikey.req_) != 0
}

func (ikey *InternalKey) Encode() Slice {
	return NewByString(ikey.req_)
}

