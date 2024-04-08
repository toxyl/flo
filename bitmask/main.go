package bitmask

import "fmt"

type Bitmask uint32

func (bm *Bitmask) Uint() uint32                      { return uint32(*bm) }
func (bm *Bitmask) String() string                    { return fmt.Sprintf("%032b", bm.Uint()) }
func (bm *Bitmask) AND(v uint32) *Bitmask             { return New(bm.Uint() & v) }
func (bm *Bitmask) AND_NOT(v uint32) *Bitmask         { return New(bm.Uint() &^ v) }
func (bm *Bitmask) OR(v uint32) *Bitmask              { return New(bm.Uint() | v) }
func (bm *Bitmask) XOR(v uint32) *Bitmask             { return New(bm.Uint() ^ v) }
func (bm *Bitmask) NOT() *Bitmask                     { return New(^bm.Uint()) }
func (bm *Bitmask) ShiftRight(n int) *Bitmask         { return New(bm.Uint() >> n) }
func (bm *Bitmask) ShiftLeft(n int) *Bitmask          { return New(bm.Uint() << n) }
func (bm *Bitmask) SetBit(n uint) *Bitmask            { return bm.OR(uint32(0b1 << n)) }
func (bm *Bitmask) ClearBit(n uint) *Bitmask          { return bm.AND_NOT(uint32(0b1 << n)) }
func (bm *Bitmask) Set(v uint32) *Bitmask             { return bm.OR(v) }
func (bm *Bitmask) Clear(v uint32) *Bitmask           { return bm.AND_NOT(v) }
func (bm *Bitmask) Match(v uint32) bool               { return (bm.Uint() & v) == v }
func (bm *Bitmask) MatchAny(v uint32) bool            { return (bm.Uint() & v) != 0 }
func (bm *Bitmask) Mask(m uint32, shift int) *Bitmask { return bm.AND(m).ShiftRight(shift) }

func New(value uint32) *Bitmask {
	b := Bitmask(value)
	return &b
}

func NewWithMask(value uint32, mask uint32, shift int) *Bitmask {
	return New(value).Mask(mask, shift)
}
