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

// dasm is the assembler for the drones VM.  Most of its functionality
// is implemented in the asm package.
package main

import (
	"bufio"
	"fmt"
	"github.com/bieber/drones/asm"
	"os"
)

func main() {
	fmt.Println(asm.Assemble(bufio.NewReader(os.Stdin)))
}
