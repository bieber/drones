package vm

// Opcode represents the executable opcodes for the VM.
type Opcode uint

const (
	_ = iota

	// Does nothing for 1 clock cycle.
	Nop Opcode = iota

	// Jumps to the argument address.
	Jmp
	// Jumps to the argument address only if a == 0.
	Jmpz

	// Loads a constant argument into a.
	Ldc
)

// OpcodeNames returns a map of opcode values to their name.
func OpcodeNames() map[Opcode]string {
	return map[Opcode]string{
		Nop:  "nop",
		Jmp:  "jmp",
		Jmpz: "jmpz",
		Ldc:  "ldc",
	}
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
	case Jmpz:
		if vm.a == 0 {
			vm.ip = arg
		}

	case Ldc:
		vm.a = arg

	default:
		panic("vm: Invalid opcode")
	}
}
