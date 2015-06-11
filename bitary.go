package bitarray

import (
	"bytes"
	"fmt"
)

//BitArray contains unexported fields
type BitArray struct {
	length int
	data   []byte
}

//New returns new instance of BitArray
func New(length int) *BitArray {
	if length < 1 {
		panic(fmt.Sprintf("pointless size of %d", length))
	}
	blen := length / 8
	if length%8 > 0 {
		blen++
	}
	d := make([]byte, blen)
	return &BitArray{length, d}
}

//Fill sets all bits in the array to filler (1 or 0)
func (r *BitArray) Fill(filler uint8) {
	if filler != 0 {
		filler = 1
	}
	for i := int(0); i < r.Length(); i++ {
		r.set(i, filler)
	}
}

//Length returns the size of the array
func (r *BitArray) Length() int {
	return r.length
}

func (r *BitArray) set(idx int, val uint8) {
	if idx < 0 || idx >= r.Length() {
		panic(fmt.Sprintf("index %d is out of bounds", idx))
	}
	bidx := idx / 8
	num := uint8(r.data[bidx])
	off := uint8(7 - (idx % 8))
	mask := uint8(1) << off
	if num&mask == mask {
		//is set
		if val == 0 {
			//so unset it
			r.data[bidx] = num ^ mask
		}
	} else {
		//not set
		if val == 1 {
			//so set it
			r.data[bidx] = num | mask
		}
	}
}

func (r *BitArray) get(idx int) uint8 {
	if idx < 0 || idx >= r.Length() {
		panic(fmt.Sprintf("index %d is out of bounds", idx))
	}
	bidx := idx / 8
	num := uint8(r.data[bidx])
	off := uint8(7 - (idx % 8))
	mask := uint8(1) << off
	if num&mask == mask {
		return 1
	}
	return 0
}

//Set sets the bit at idx to 1
func (r *BitArray) Set(idx int) {
	r.set(idx, 1)
}

//Unset sets the bit at idx to 0
func (r *BitArray) Unset(idx int) {
	r.set(idx, 0)
}

//IsSet returns true if bit at idx is set to 1
func (r *BitArray) IsSet(idx int) bool {
	return r.get(idx) == 1
}

//IsUnset returns true if bit at idx is set to 0
func (r *BitArray) IsUnset(idx int) bool {
	return r.get(idx) == 0
}

//String returns text representation of the array
func (r *BitArray) String() string {
	var buf bytes.Buffer
	for _, v := range r.data {
		buf.WriteString(fmt.Sprintf("%08b", v))
	}
	return buf.String()[0:r.length]
}
