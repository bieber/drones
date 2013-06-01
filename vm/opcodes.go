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

// Opcode represents the executable opcodes for the VM.
type Opcode uint

const (
	_ = iota

	// Does nothing for 1 clock cycle.
	Nop Opcode = iota

	// Jumps to the argument address.
	Jmp
	// Jumps to the argument address only if a == 0.
	Jz
	// Jumps to the argument address only if a != 0.
	Jnz

	// Swaps the a and b registers.
	Sab
	// Swaps the a and p registers.
	Sap
	// Swaps the a and i registers.
	Sai
	// Loads the b register into p.
	Lbp

	// Loads a constant argument into a.
	Ldc
	// Loads from memory at the argument address into a.
	Ldm
	// Loads from memory at the address in p into a.
	Ldp
	// Loads from memory at the address in (p + i) into a.
	Ldi

	// Pushes the content of a onto the stack.
	Push
	// Pops the top of the stack into a.
	Pop

	// Reads from the bus number in the argument into a.
	Rb
	// Writes from a into the bus number in the argument.
	Wb

	// Jumps to the address in the argument preparing a new stack
	// frame for a function call.  Does not automatically preserve
	// registers, push them manually if you want to save them.
	Call
	// Returns from a function called with call.
	Ret
)

// OpcodeNames returns a map of opcode values to their name.
func OpcodeNames() map[Opcode]string {
	return map[Opcode]string{
		Nop:  "nop",
		Jmp:  "jmp",
		Jz:   "jz",
		Jnz:  "jnz",
		Sab:  "sab",
		Sap:  "sap",
		Sai:  "sai",
		Lbp:  "lbp",
		Ldc:  "ldc",
		Ldm:  "ldm",
		Ldp:  "ldp",
		Ldi:  "ldi",
		Push: "push",
		Pop:  "pop",
		Rb:   "rb",
		Wb:   "wb",
		Call: "call",
		Ret:  "ret",
	}
}

// ClockN runs Clock() on the vm n times.
func (vm *VM) ClockN(n int) {
	for i := 0; i < n; i++ {
		vm.Clock()
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

		// Jumps
	case Jmp:
		vm.ip = arg
	case Jz:
		if vm.a == 0 {
			vm.ip = arg
		}
	case Jnz:
		if vm.a != 0 {
			vm.ip = arg
		}

		// Swaps
	case Sab:
		vm.b = vm.a ^ vm.b
		vm.a = vm.a ^ vm.b
		vm.b = vm.a ^ vm.b
	case Sap:
		vm.p = vm.a ^ vm.p
		vm.a = vm.a ^ vm.p
		vm.p = vm.a ^ vm.p
	case Sai:
		vm.i = vm.a ^ vm.i
		vm.a = vm.a ^ vm.i
		vm.i = vm.a ^ vm.i
	case Lbp:
		vm.p = vm.bp

		// Loads
	case Ldc:
		vm.a = arg
	case Ldm:
		vm.a = vm.mem[arg]
	case Ldp:
		vm.a = vm.mem[vm.p]
	case Ldi:
		vm.a = vm.mem[vm.p+vm.i]

		// Stack manipulation
	case Push:
		vm.push(vm.a)
	case Pop:
		vm.a = vm.pop()

		// Bus communication
	case Rb:
		vm.a = vm.Buses[arg]
	case Wb:
		vm.Buses[arg] = vm.a

		// Function calls
	case Call:
		newBase := vm.sp
		vm.push(vm.ip)
		vm.push(vm.bp)
		vm.bp = newBase
		vm.ip = arg
	case Ret:
		vm.sp = vm.bp - 2
		vm.bp = vm.pop()
		vm.ip = vm.pop()

	default:
		panic("vm: Invalid opcode")
	}
}

func (vm *VM) push(value uint16) {
	vm.mem[vm.sp] = value
	vm.sp--
}

func (vm *VM) pop() (value uint16) {
	value = vm.mem[vm.sp+1]
	vm.sp++
	return
}