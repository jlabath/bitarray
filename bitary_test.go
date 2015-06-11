package bitarray

import (
	"math"
	"testing"
)

func TestBasic(t *testing.T) {
	d := New(8)
	equals(t, d.Length(), 8)
	assert(t, d.IsUnset(0), "expected first bit unset")
	d.Set(0)
	assert(t, d.IsSet(0), "expected first bit set")
	assert(t, d.IsUnset(1), "expected bit unset")
	d.Set(1)
	assert(t, d.IsSet(1), "expected bit set")
	equals(t, "11000000", d.String())
	d.Set(4)
	d.Set(5)
	d.Set(6)
	d.Set(7)
	equals(t, "11001111", d.String())
	assert(t, d.IsUnset(2), "expected bit unset")
	assert(t, d.IsUnset(3), "expected bit unset")
	assert(t, d.IsSet(4), "expected bit set")
	assert(t, d.IsSet(5), "expected bit set")
	assert(t, d.IsSet(6), "expected bit set")
	assert(t, d.IsSet(7), "expected bit set")
	d.Unset(7)
	assert(t, d.IsUnset(7), "expected bit unset")
	d.Unset(0)
	assert(t, d.IsUnset(0), "expected bit unset")
}

func TestOddSize(t *testing.T) {
	d := New(19)
	equals(t, d.Length(), 19)
	equals(t, d.String(), "0000000000000000000")
	assert(t, d.IsUnset(18), "expected bit unset")
	assert(t, d.IsUnset(10), "expected bit unset")
	assert(t, d.IsUnset(1), "expected bit unset")
	d.Set(18)
	equals(t, d.String(), "0000000000000000001")
	assert(t, d.IsSet(18), "expected bit set")
	d.Set(7)
	assert(t, d.IsSet(7), "expected bit set")
	d.Set(15)
	assert(t, d.IsSet(15), "expected bit set")
	equals(t, d.String(), "0000000100000001001")
}

func TestFill(t *testing.T) {
	var l int = 0xFFFFFF
	d := New(l)
	equals(t, d.Length(), l)
	for i := 0; i < d.Length(); i++ {
		assert(t, d.IsUnset(i), "expected bit unset")
	}
	d.Fill(1)
	for i := 0; i < d.Length(); i++ {
		assert(t, d.IsSet(i), "expected bit set")
	}
	d.Fill(0)
	for i := 0; i < d.Length(); i++ {
		assert(t, d.IsUnset(i), "expected bit unset")
	}
}

func TestBig(t *testing.T) {
	//may crash if not enough memory as make checks availabale memory
	l := math.MaxInt32
	d := New(l)
	equals(t, d.Length(), l)
	idx := l - 4
	assert(t, d.IsUnset(idx), "expected bit unset")
	d.Set(idx)
	assert(t, d.IsSet(idx), "expected bit set")
	d.Unset(idx)
	assert(t, d.IsUnset(idx), "expected bit unset")
}

func TestAbuse(t *testing.T) {
	checkPanic := func(f func(), name string) {
		defer func() {
			r := recover()
			assert(t, r != nil, "call %s was supposed to panic", name)
		}()
		f()
	}
	f1 := func() {
		d := New(-1)
		_ = d
	}
	checkPanic(f1, "f1")
	f2 := func() {
		d := New(1)
		d.IsSet(1)
	}
	checkPanic(f2, "f2")
	f3 := func() {
		d := New(1)
		d.IsSet(-1)
	}
	checkPanic(f3, "f3")
	f4 := func() {
		d := New(1)
		d.Set(1)
	}
	checkPanic(f4, "f4")
	f5 := func() {
		d := New(1)
		d.Set(-1)
	}
	checkPanic(f5, "f5")
	f6 := func() {
		d := New(1)
		d.Unset(1)
	}
	checkPanic(f6, "f6")
	f7 := func() {
		d := New(1)
		d.IsUnset(1)
	}
	checkPanic(f7, "f7")
}
