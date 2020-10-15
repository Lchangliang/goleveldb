package util

type Comparator interface {
	Compare(a []byte, b []byte) int 

}