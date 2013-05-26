// Package vm implements the drones virtual machine, a 16 bit VM intended to
// drive fictional robots.
package vm

import (
	"strings"
	"fmt"
)

/*
 VM is an instance of the virtual machine.  It stores the VM's current state,
 and its methods can be used to mutate its state.  The buses are directly
 addressable, and allow the VM to communicate with external components.
*/
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

// New allocates and initializes memory, buses and registers for a new VM.
func New(memSize uint16, buses uint16) (vm *VM) {
	vm = &VM{MemSize: memSize, sp: memSize - 1, bp: memSize - 1}
	vm.mem = make([]uint16, memSize)
	vm.Buses = make([]uint16, buses)
	return
}

// Debug returns a string that fully specifies the current state of the VM.
func (vm *VM) Debug() string {
	titleLine := fmt.Sprintf("%7s %8s %8s %6s", "", "Signed", "Unsigned", "Hex")
	lines := []string{
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
	}
	for i, v := range vm.mem {
		lines = append(lines, debugLine(fmt.Sprintf("%#04x", i), v))
	}
	lines = append(lines, "\n=== Buses ===\n")
	for i, v := range vm.Buses {
		lines = append(lines, debugLine(fmt.Sprintf("%#04x", i), v))
	}
	return strings.Join(lines, "\n")
}

// Generates a line of the VM debug output for a given value and label
func debugLine(label string, value uint16) string {
	return fmt.Sprintf("%6s: %8d %8d %#04x", label, value, int16(value), value)
}