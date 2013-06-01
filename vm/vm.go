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

// Package vm implements the drones virtual machine, a 16 bit VM intended to
// drive fictional robots.
package vm

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

// VM is an instance of the virtual machine.  It stores the VM's
// current state, and its methods can be used to mutate its state.
// The buses are directly addressable, and allow the VM to communicate
// with external components.
type VM struct {
	// Registers
	ip uint16
	bp uint16
	sp uint16
	a  uint16
	b  uint16
	p  uint16
	i  uint16

	// The number of 16-bit words in memory.
	MemSize uint16
	// Main memory.
	mem []uint16

	// Externally accessible buses.
	Buses []uint16
}

// New allocates and initializes memory, buses and registers for a new
// VM.
func New(memSize uint16, buses uint16) (vm *VM) {
	vm = &VM{MemSize: memSize, sp: memSize - 1, bp: memSize - 1}
	vm.mem = make([]uint16, memSize)
	vm.Buses = make([]uint16, buses)
	return
}

// LoadMem copies a slice of uint16 into the VM's memory.
func (vm *VM) LoadMem(memory []uint16) {
	copy(vm.mem, memory)
}

// LoadOpcodes copies a slice of Opcodes into the VM's memory.
func (vm *VM) LoadOpcodes(memory []Opcode) {
	for i, v := range memory {
		if i >= len(vm.mem) {
			return
		}
		vm.mem[i] = uint16(v)
	}
}

// LoadBinary reads a binary data file into the VM's memory.  The
// binary file should begin with a single 16-bit value designating the
// number of 16-bit words in the file, followed by that many values.
// Byte-order is expected to be little-endian.
func (vm *VM) LoadBinary(r io.Reader) error {
	var size uint16
	binary.Read(r, binary.LittleEndian, &size)
	if int(size) < len(vm.mem) {
		vm.mem = vm.mem[:size]
	}
	err := binary.Read(r, binary.LittleEndian, vm.mem)
	vm.mem = vm.mem[:cap(vm.mem)]
	return err
}

// Debug returns a string that fully specifies the current state of
// the VM.
func (vm *VM) Debug() string {
	titleLine := fmt.Sprintf(
		"%6s  %6s %8s %8s %6s",
		"",
		"Opcode",
		"Signed",
		"Unsigned",
		"Hex",
	)
	lines := []string{
		"=================",
		"=== Registers ===\n",
		titleLine,
		debugLine("ip", vm.ip),
		debugLine("bp", vm.bp),
		debugLine("sp", vm.sp),
		debugLine("a", vm.a),
		debugLine("b", vm.b),
		debugLine("p", vm.p),
		debugLine("i", vm.i),
		"\n=== Memory ===\n",
		titleLine,
	}
	for i, v := range vm.mem {
		lines = append(lines, debugLine(fmt.Sprintf("%#04x", i), v))
	}
	lines = append(lines, "\n=== Buses ===\n", titleLine)
	for i, v := range vm.Buses {
		lines = append(lines, debugLine(fmt.Sprintf("%#04x", i), v))
	}
	lines = append(lines, "\n")
	return strings.Join(lines, "\n")
}

// Generates a line of the VM debug output for a given value and label
func debugLine(label string, value uint16) string {
	opcodeName, ok := OpcodeNames()[Opcode(value)]
	if !ok {
		opcodeName = "n/a"
	}
	return fmt.Sprintf(
		"%6s: %6s %8d %8d %#04x",
		label,
		opcodeName,
		int16(value),
		value,
		value,
	)
}
