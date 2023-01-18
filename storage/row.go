package storage

import (
	"encoding/binary"
	"math"
)

type row struct {
	data []byte
	col  *column
}

func (r *row) getLen() int {
	return len(r.data)
}

func (r *row) setData(t []byte) {
	r.data = t
}

// return the size of ith column
func (r *row) getIthSize(i int) int {
	b := r.data[i*4 : (i+1)*4]
	return int(binary.BigEndian.Uint32(b))
}

// return the ith value of table
func (r *row) getIth(i, offset int) interface{} {
	size := r.getIthSize(i)
	n := len(r.data)
	data := r.data[n-offset-size : n-offset]
	t := r.col.getIth(i)
	switch t {
	case IntType:
		return binary.BigEndian.Uint32(data)
	case FloatType:
		return math.Float32frombits(binary.BigEndian.Uint32(data))
	case StrType:
	default:
		return string(data)
	}
	return nil
}

// GetRow return a new row of given data
func GetRow(data []byte) *row {
	r := row{}
	r.setData(data)
	return &r
}
