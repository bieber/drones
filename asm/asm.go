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

// Package asm implements the internals of the drones VM assembler.
package asm

import (
	_ "fmt"
)

var errorCodes []string
var code []uint16 = []uint16{1, 2, 3}

func addLabel(label string) {
}

func setOrg(addr uint16) {
}

func addOpcodeLabel(opcode string, label string) {

}

func addOpcodeConstant(opcode string, constant uint16) {

}
