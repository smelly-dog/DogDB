package storage

import "fmt"

const (
	IntType uint8 = iota
	FloatType
	StrType
)

// describe the type of columns in the table
type column struct {
	tyData []uint8
}

// set the type of ith column
func (c *column) setIth(i int, t byte) {
	if i >= len(c.tyData) || i < 0 {
		msg := fmt.Sprintf("the index is out of bounds, [index->%d], [len->%d]", i, len(c.tyData))
		panic(msg)
	}
	c.tyData[i] = t
}

// return the type of ith column
func (c *column) getIth(i int) uint8 {
	if i >= len(c.tyData) || i < 0 {
		msg := fmt.Sprintf("the index is out of bounds, [index->%d], [len->%d]", i, len(c.tyData))
		panic(msg)
	}
	return c.tyData[i]
}

// GetNewColumn return a new column description
func GetNewColumn(b []byte) *column {
	col := column{
		tyData: make([]uint8, len(b), len(b)),
	}
	for i := 0; i < len(b); i++ {
		col.setIth(i, b[i])
	}
	return &col
}
