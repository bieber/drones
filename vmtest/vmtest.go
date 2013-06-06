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

// vmtest runs test binaries through the VM and debugs along the way.
package main

import (
	"encoding/binary"
	"fmt"
	"github.com/bieber/drones/vm"
)

func main() {
	v := vm.New(200, 3)
	v.LoadOpcodes(
		[]vm.Opcode{
			vm.Ldm, 52,
			vm.Wb, 0,
			vm.Ldc, 1,
			vm.Wb, 1,
			vm.Ldm, 53,
			vm.Wb, 0,
			vm.Ldc, 1,
			vm.Wb, 1,
			vm.Ldm, 54,
			vm.Wb, 0,
			vm.Ldc, 1,
			vm.Wb, 1,
			vm.Ldm, 55,
			vm.Wb, 0,
			vm.Ldc, 1,
			vm.Wb, 1,
			vm.Ldm, 56,
			vm.Wb, 0,
			vm.Ldc, 1,
			vm.Wb, 1,
			vm.Ldm, 57,
			vm.Wb, 0,
			vm.Ldc, 1,
			vm.Wb, 1,
			vm.Ldc, 1,
			vm.Wb, 2,
			vm.Opcode(binary.LittleEndian.Uint16([]byte("He"))),
			vm.Opcode(binary.LittleEndian.Uint16([]byte("ll"))),
			vm.Opcode(binary.LittleEndian.Uint16([]byte("o "))),
			vm.Opcode(binary.LittleEndian.Uint16([]byte("Wo"))),
			vm.Opcode(binary.LittleEndian.Uint16([]byte("rl"))),
			vm.Opcode(binary.LittleEndian.Uint16([]byte("d\n"))),
		},
	)
	for {
		v.Clock()
		if v.Buses[1] != 0 {
			v.Buses[1] = 0
			chars := make([]byte, 2)
			binary.LittleEndian.PutUint16(chars, v.Buses[0])
			fmt.Print(string(chars))
		}
		if v.Buses[2] != 0 {
			break
		}
	}
}
