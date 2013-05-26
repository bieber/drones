// Package vm implements the drones virtual machine, a 16 bit VM intended to
// drive fictional robots.
package vm

/*
 VM is an instance of the virtual machine.  It contains the VM's memory
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

	// Memory
	MemSize uint16
	mem []uint16

	// Externally accessible buses.
	Buses []uint16
}

func New(memSize uint16, buses uint16) (vm *VM) {
	vm = &VM{MemSize: memSize, sp: memSize - 1, bp: memSize - 1}
	vm.mem = make([]uint16, memSize)
	vm.Buses = make([]uint16, buses)
	return
}
