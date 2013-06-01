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
	"bytes"
	"fmt"
	"github.com/bieber/drones/vm"
)

func main() {
	mem := []byte{5, 0, 4, 0, 5, 3, 6, 25, 78, 12, 3, 4}
	v := vm.New(4, 5)
	v.LoadBinary(bytes.NewReader(mem))
	fmt.Println(v.Debug())
	v.Clock()
	fmt.Println(v.Debug())
}
