/*
 * Copyright 2013, Robert Bieber
 *
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

package asm

import (
	"bufio"
	"github.com/bieber/drones/vm"
	. "launchpad.net/gocheck"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type suite struct{}

var _ = Suite(&suite{})

func (s *suite) TestOpcodes(c *C) {
	code := ""
	opcodes := make([]uint16, 0, len(vm.OpcodeNames()))
	for opcode, name := range vm.OpcodeNames() {
		code = code + name + "\n"
		opcodes = append(opcodes, uint16(opcode))
	}
	binary, err := Assemble(bufio.NewReader(strings.NewReader(code)))

	c.Assert(err, IsNil)
	c.Assert(binary, NotNil)
	c.Assert(len(binary)/2, Equals, len(vm.OpcodeNames()))

	i := 0
	for _, opcode := range opcodes {
		c.Assert(binary[i], Equals, opcode)
		i += 2
	}
}

func (s *suite) TestLabels(c *C) {
	code := `
        nop
label1: nop
        nop
        nop
label2:
        nop
%org 20
        nop label1
label3: jmp label2
        jmp label3

`
	binary, err := Assemble(bufio.NewReader(strings.NewReader(code)))

	c.Assert(err, IsNil)
	c.Assert(binary, NotNil)
	c.Assert(binary[21], Equals, uint16(2))
	c.Assert(binary[23], Equals, uint16(8))
	c.Assert(binary[25], Equals, uint16(22))
}

func (s *suite) TestWords(c *C) {
	code := `
%words 50 0x10
%org 20
%words 10
`
	binary, err := Assemble(bufio.NewReader(strings.NewReader(code)))

	c.Assert(err, IsNil)
	c.Assert(binary, NotNil)
	c.Assert(binary[0], Equals, uint16(50))
	c.Assert(binary[1], Equals, uint16(16))
	c.Assert(binary[20], Equals, uint16(10))
}
