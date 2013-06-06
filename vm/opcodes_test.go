/*
 * This file is part of drones.
 *
 * drones is free software: you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * drones is distributed in the hope that it will be useful,
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with drones.  If not, see <http://www.gnu.org/licenses/>.
 */

package vm

import (
	"encoding/binary"
	ck "launchpad.net/gocheck"
	"math/rand"
	"testing"
)

func Test(t *testing.T) {
	ck.TestingT(t)
}

type suite struct {
	vm *VM
	r  *rand.Rand
}

var _ = ck.Suite(&suite{})

func (s *suite) SetUpTest(c *ck.C) {
	s.vm = New(100, 10)
	// Using the same seed every time keeps the random tests deterministic
	s.r = rand.New(rand.NewSource(50))
}

func (s *suite) TestJmp(c *ck.C) {
	s.vm.LoadOpcodes(
		[]Opcode{
			Jmp, 6,
		},
	)
	s.vm.Clock()
	c.Assert(s.vm.ip, ck.Equals, uint16(6))
}

func (s *suite) TestSwaps(c *ck.C) {
	s.vm.LoadOpcodes(
		[]Opcode{
			Ldc, 1,
			Sab, 0,
			Ldc, 2,
			Sap, 0,
			Ldc, 3,
			Sai, 0,
			Lbp, 0,
		},
	)

	s.vm.ClockN(2)
	c.Assert(s.vm.b, ck.Equals, uint16(1))
	c.Assert(s.vm.a, ck.Equals, uint16(0))

	s.vm.ClockN(2)
	c.Assert(s.vm.p, ck.Equals, uint16(2))
	c.Assert(s.vm.a, ck.Equals, uint16(0))

	s.vm.ClockN(2)
	c.Assert(s.vm.i, ck.Equals, uint16(3))
	c.Assert(s.vm.a, ck.Equals, uint16(0))

	s.vm.Clock()
	c.Assert(s.vm.p, ck.Equals, uint16(len(s.vm.mem)-1))
	c.Assert(s.vm.bp, ck.Equals, s.vm.p)
}

func (s *suite) TestJumps(c *ck.C) {
	s.vm.LoadOpcodes(
		[]Opcode{
			Jz, 6,
			Ldc, 0,
			Jnz, 0,
			Ldc, 2,
			Jz, 0,
			Jmp, 14,
			Ldc, 5,
			Jnz, 2,
		},
	)

	s.vm.Clock()
	c.Assert(s.vm.ip, ck.Equals, uint16(6))

	s.vm.Clock()
	c.Assert(s.vm.ip, ck.Equals, uint16(8))
	c.Assert(s.vm.a, ck.Equals, uint16(2))

	s.vm.Clock()
	c.Assert(s.vm.ip, ck.Equals, uint16(10))

	s.vm.Clock()
	c.Assert(s.vm.ip, ck.Equals, uint16(14))

	s.vm.Clock()
	c.Assert(s.vm.ip, ck.Equals, uint16(2))

	s.vm.Clock()
	c.Assert(s.vm.ip, ck.Equals, uint16(4))
	c.Assert(s.vm.a, ck.Equals, uint16(0))

	s.vm.Clock()
	c.Assert(s.vm.ip, ck.Equals, uint16(6))
}

func (s *suite) TestLoads(c *ck.C) {
	s.vm.LoadOpcodes(
		[]Opcode{
			Ldc, 3,
			Sap, 500,
			Ldc, 4,
			Sai, 700,
			Ldp, 900,
			Ldi, 0,
			Ldm, 9,
		},
	)

	s.vm.Clock()
	c.Assert(s.vm.a, ck.Equals, uint16(3))

	s.vm.ClockN(4)
	c.Assert(s.vm.a, ck.Equals, uint16(500))

	s.vm.Clock()
	c.Assert(s.vm.a, ck.Equals, uint16(700))

	s.vm.Clock()
	c.Assert(s.vm.a, ck.Equals, uint16(900))
}

func (s *suite) TestStack(c *ck.C) {
	s.vm.LoadOpcodes(
		[]Opcode{
			Ldc, 10,
			Push, 0,
			Ldc, 20,
			Push, 0,
			Ldc, 30,
			Pop, 0,
			Ldc, 40,
			Pop, 0,
		},
	)

	top := uint16(len(s.vm.mem) - 1)

	s.vm.ClockN(2)
	c.Assert(s.vm.sp, ck.Equals, top-1)
	c.Assert(s.vm.mem[top], ck.Equals, uint16(10))

	s.vm.ClockN(2)
	c.Assert(s.vm.sp, ck.Equals, top-2)
	c.Assert(s.vm.mem[top-1], ck.Equals, uint16(20))

	s.vm.ClockN(2)
	c.Assert(s.vm.sp, ck.Equals, top-1)
	c.Assert(s.vm.a, ck.Equals, uint16(20))

	s.vm.ClockN(2)
	c.Assert(s.vm.sp, ck.Equals, top)
	c.Assert(s.vm.a, ck.Equals, uint16(10))
}

func (s *suite) TestBuses(c *ck.C) {
	s.vm.LoadOpcodes(
		[]Opcode{
			Ldc, 5,
			Wb, 0,
			Rb, 1,
			Wb, 2,
		},
	)

	s.vm.ClockN(2)
	c.Assert(s.vm.Buses[0], ck.Equals, uint16(5))

	s.vm.Buses[1] = 15
	s.vm.Clock()
	c.Assert(s.vm.a, ck.Equals, uint16(15))

	s.vm.Clock()
	c.Assert(s.vm.Buses[2], ck.Equals, uint16(15))
}

func (s *suite) TestFuncs(c *ck.C) {
	s.vm.LoadOpcodes(
		[]Opcode{
			Ldc, 29,
			Push, 0,
			Call, 12,
			Jmp, 100,
			Ldc, 5, // Second function
			Ret, 0,
			Lbp, 0, // First function
			Ldc, 1,
			Sai, 0,
			Ldi, 0,
			Call, 8,
			Ret, 0,
		},
	)

	top := uint16(len(s.vm.mem) - 1)

	s.vm.ClockN(3)
	c.Assert(s.vm.ip, ck.Equals, uint16(12))
	c.Assert(s.vm.mem[s.vm.bp], ck.Equals, uint16(6))
	c.Assert(s.vm.mem[s.vm.bp-1], ck.Equals, top)
	f1bp := s.vm.bp

	s.vm.ClockN(4)
	c.Assert(s.vm.a, ck.Equals, uint16(29))

	s.vm.Clock()
	c.Assert(s.vm.ip, ck.Equals, uint16(8))
	c.Assert(s.vm.mem[s.vm.bp], ck.Equals, uint16(22))
	c.Assert(s.vm.mem[s.vm.bp-1], ck.Equals, f1bp)

	s.vm.Clock()
	c.Assert(s.vm.a, ck.Equals, uint16(5))

	s.vm.Clock()
	c.Assert(s.vm.ip, ck.Equals, uint16(22))
	c.Assert(s.vm.mem[s.vm.bp], ck.Equals, uint16(6))
	c.Assert(s.vm.mem[s.vm.bp-1], ck.Equals, top)

	s.vm.Clock()
	c.Assert(s.vm.ip, ck.Equals, uint16(6))
	c.Assert(s.vm.bp, ck.Equals, top)

	s.vm.Clock()
	c.Assert(s.vm.ip, ck.Equals, uint16(100))
}

func (s *suite) TestArithmetic(c *ck.C) {
	s.vm.LoadOpcodes(
		[]Opcode{
			Ldc, 1,
			Sab, 0,
			Ldc, 5,
			Add, 0,
			Sab, 0,
			Ldc, 13,
			Sub, 0,
			Sab, 0,
			Ldc, 2,
			Mul, 0,
			Sab, 0,
			Ldc, 0xfffe,
			Mul, 0,
			Sab, 0,
			Ldc, 0xfffe,
			Sab, 0,
			Sdiv, 0,
			Sab, 0,
			Ldc, 8,
			Sab, 0,
			Div, 0,
		},
	)

	s.vm.ClockN(4)
	c.Assert(s.vm.a, ck.Equals, uint16(6))
	s.vm.ClockN(3)
	c.Assert(s.vm.a, ck.Equals, uint16(7))
	s.vm.ClockN(3)
	c.Assert(s.vm.a, ck.Equals, uint16(14))
	s.vm.ClockN(3)
	c.Assert(s.vm.a, ck.Equals, uint16(0xffe4))
	s.vm.ClockN(4)
	c.Assert(s.vm.a, ck.Equals, uint16(14))
	s.vm.ClockN(4)
	c.Assert(s.vm.a, ck.Equals, uint16(1))
	c.Assert(s.vm.b, ck.Equals, uint16(6))
}

func (s *suite) TestBitwise(c *ck.C) {
	s.vm.LoadOpcodes(
		[]Opcode{
			Ldc, 0x6666,
			Sab, 0,
			Ldc, 0x3333,
			And, 0,
			Sab, 0,
			Ldc, 0x1111,
			Or, 0,
			Sab, 0,
			Ldc, 0x2222,
			Xor, 0,
			Sab, 0,
			Ldc, 0x2222,
			Xor, 0,
			Sab, 0,
			Ldc, 2,
			Sab, 0,
			Shl, 0,
			Sab, 0,
			Ldc, 1,
			Sab, 0,
			Shr, 0,
			Not, 0,
		},
	)
	s.vm.ClockN(4)
	c.Assert(s.vm.a, ck.Equals, uint16(0x2222))
	s.vm.ClockN(3)
	c.Assert(s.vm.a, ck.Equals, uint16(0x3333))
	s.vm.ClockN(3)
	c.Assert(s.vm.a, ck.Equals, uint16(0x1111))
	s.vm.ClockN(3)
	c.Assert(s.vm.a, ck.Equals, uint16(0x3333))
	s.vm.ClockN(4)
	c.Assert(s.vm.a, ck.Equals, uint16(0xcccc))
	s.vm.ClockN(4)
	c.Assert(s.vm.a, ck.Equals, uint16(0x6666))
	s.vm.Clock()
	c.Assert(s.vm.a, ck.Equals, uint16(0x9999))
}

func (s *suite) TestComparisons(c *ck.C) {
	opcodes := []Opcode{
		Ldc, 0,
		Sab, 0,
		Ldc, 0,
		0, 0,
	}

	for i := 0; i < 100; i++ {
		a, b := s.randPair()
		if i%20 == 0 {
			// Otherwise won't get any a == b cases
			a = b
		}
		opcodes[5] = Opcode(a)
		opcodes[1] = Opcode(b)

		opcodes[6] = Lt
		s.reset(opcodes)
		s.vm.ClockN(4)
		c.Assert(s.vm.a, ck.Equals, vmBool(a < b))

		opcodes[6] = Lts
		s.reset(opcodes)
		s.vm.ClockN(4)
		c.Assert(s.vm.a, ck.Equals, vmBool(int16(a) < int16(b)))

		opcodes[6] = Le
		s.reset(opcodes)
		s.vm.ClockN(4)
		c.Assert(s.vm.a, ck.Equals, vmBool(a <= b))

		opcodes[6] = Les
		s.reset(opcodes)
		s.vm.ClockN(4)
		c.Assert(s.vm.a, ck.Equals, vmBool(int16(a) <= int16(b)))

		opcodes[6] = Gt
		s.reset(opcodes)
		s.vm.ClockN(4)
		c.Assert(s.vm.a, ck.Equals, vmBool(a > b))

		opcodes[6] = Gts
		s.reset(opcodes)
		s.vm.ClockN(4)
		c.Assert(s.vm.a, ck.Equals, vmBool(int16(a) > int16(b)))

		opcodes[6] = Ge
		s.reset(opcodes)
		s.vm.ClockN(4)
		c.Assert(s.vm.a, ck.Equals, vmBool(a >= b))

		opcodes[6] = Ges
		s.reset(opcodes)
		s.vm.ClockN(4)
		c.Assert(s.vm.a, ck.Equals, vmBool(int16(a) >= int16(b)))

		opcodes[6] = Eq
		s.reset(opcodes)
		s.vm.ClockN(4)
		c.Assert(s.vm.a, ck.Equals, vmBool(a == b))

		opcodes[6] = Neq
		s.reset(opcodes)
		s.vm.ClockN(4)
		c.Assert(s.vm.a, ck.Equals, vmBool(a != b))
	}
}

func (s *suite) randPair() (a uint16, b uint16) {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, s.r.Uint32())
	a = binary.LittleEndian.Uint16(data[0:2])
	b = binary.LittleEndian.Uint16(data[2:4])
	return
}

func (s *suite) reset(mem []Opcode) {
	s.vm.LoadOpcodes(mem)
	s.vm.ip = 0
}
