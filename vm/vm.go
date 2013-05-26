// Package vm implements the drones virtual machine, a 16 bit VM intended to
// drive fictional robots.
package vm

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

// Opcode represents the executable opcodes for the VM.
type Opcode uint

const (
	_ = iota

	// Does nothing for 1 clock cycle.
	Nop Opcode = iota

	// Jumps to the argument address.
	Jmp
	// Jumps to the argument address only if a != 0.
	Jmpc

	// Loads a constant argument into a.
	Ldc
)

// OpcodeNames returns a map of opcode values to their name.
func OpcodeNames() map[Opcode]string {
	return map[Opcode]string{
		Nop:  "nop",
		Jmp:  "jmp",
		Jmpc: "jmpc",
		Ldc:  "ldc",
	}
}

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

// Clock advances the VM by one clock cycle.  The opcode located at
// the memory location at index ip will be executed with the opcode
// located at memory location ip + 1 used as an argument.  See the
// Opcode constants for more information.
func (vm *VM) Clock() {
	// If something goes wrong reset ip to 0 to simulate a machine
	// reset.
	defer func() {
		if recover() != nil {
			vm.ip = 0
		}
	}()

	opcode := vm.mem[vm.ip]
	arg := vm.mem[vm.ip+1]
	vm.ip += 2
	switch Opcode(opcode) {
	case Nop:
		// Do nothing

	case Jmp:
		vm.ip = arg
	case Jmpc:
		if vm.a != 0 {
			vm.ip = arg
		}

	case Ldc:
		vm.a = arg

	default:
		panic("vm: Invalid opcode")
	}
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
		value,
		int16(value),
		value,
	)
}
