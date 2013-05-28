package vm

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type suite struct {
	vm *VM
}

var _ = Suite(&suite{})

func (s *suite) SetUpTest(c *C) {
	s.vm = New(20, 2)
}

func (s *suite) TestJmp(c *C) {
	s.vm.LoadOpcodes(
		[]Opcode{
			Jmp, 6,
		},
	)
	s.vm.Clock()
	c.Assert(s.vm.ip, Equals, uint16(6))
}

func (s *suite) TestSwaps(c *C) {
	s.vm.LoadOpcodes(
		[]Opcode{
			Ldc, 1,
			Sab, 0,
			Ldc, 2,
			Sap, 0,
			Ldc, 3,
			Sai, 0,
		},
	)

	s.vm.Clock()
	s.vm.Clock()
	c.Assert(s.vm.b, Equals, uint16(1))

	s.vm.Clock()
	s.vm.Clock()
	c.Assert(s.vm.p, Equals, uint16(2))

	s.vm.Clock()
	s.vm.Clock()
	c.Assert(s.vm.i, Equals, uint16(3))
}

func (s *suite) TestJumps(c *C) {
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
	c.Assert(s.vm.ip, Equals, uint16(6))

	s.vm.Clock()
	c.Assert(s.vm.ip, Equals, uint16(8))
	c.Assert(s.vm.a, Equals, uint16(2))

	s.vm.Clock()
	c.Assert(s.vm.ip, Equals, uint16(10))

	s.vm.Clock()
	c.Assert(s.vm.ip, Equals, uint16(14))

	s.vm.Clock()
	c.Assert(s.vm.ip, Equals, uint16(2))

	s.vm.Clock()
	c.Assert(s.vm.ip, Equals, uint16(4))
	c.Assert(s.vm.a, Equals, uint16(0))

	s.vm.Clock()
	c.Assert(s.vm.ip, Equals, uint16(6))
}

func (s *suite) TestLoads(c *C) {
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
	c.Assert(s.vm.a, Equals, uint16(3))

	s.vm.Clock()
	s.vm.Clock()
	s.vm.Clock()
	s.vm.Clock()
	c.Assert(s.vm.a, Equals, uint16(500))

	s.vm.Clock()
	c.Assert(s.vm.a, Equals, uint16(700))

	s.vm.Clock()
	c.Assert(s.vm.a, Equals, uint16(900))
}
