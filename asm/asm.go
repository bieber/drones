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
	"bufio"
	"errors"
	"fmt"
	"github.com/bieber/drones/vm"
	"strings"
)

var lex *lexer
var errorCodes []string = nil

type labelInfo struct {
	address uint16
	line    int
}

type wordInfo struct {
	words   []uint16
	address uint16
	line    int
}

type opInfo struct {
	op      string
	address uint16
	label   string
	num     uint16
	line    int
}

var address uint16 = 0
var maxAddress uint16 = 0
var labels map[string]labelInfo = make(map[string]labelInfo)
var words []wordInfo
var ops []opInfo

// Assemble parses an assembly language program from in and returns the
// assembled binary, or an error if something goes wrong.
func Assemble(in *bufio.Reader) ([]uint16, error) {
	lex = newLexer(in)
	if parse, code := yyParse(lex), buildCode(); parse == 0 && code != nil {
		return code, nil
	} else {
		return nil, errors.New(strings.Join(errorCodes, "\n"))
	}
}

func buildCode() []uint16 {
	if errorCodes != nil {
		return nil
	}

	opcodeNames := vm.OpcodeNames()
	opcodes := make(map[string]uint16, len(opcodeNames))
	for k, v := range opcodeNames {
		opcodes[v] = uint16(k)
	}

	mem := make([]uint16, maxAddress)

	for _, op := range ops {
		code, ok := opcodes[op.op]
		if !ok {
			errorCodes = append(
				errorCodes,
				fmt.Sprintf("Line %d: Invalid opcode %s", op.line, op.op),
			)
			return nil
		}

		var arg uint16
		if op.label == "" {
			arg = op.num
		} else {
			label, ok := labels[op.label]
			if !ok {
				errorCodes = append(
					errorCodes,
					fmt.Sprintf("Line %d: Invalid label %s", op.line, op.label),
				)
				return nil
			}
			arg = label.address
		}
		mem[op.address] = code
		mem[op.address+1] = arg
	}

	for _, w := range words {
		for k, v := range w.words {
			mem[int(w.address)+k] = v
		}
	}

	return mem
}

func addLabel(label string) {
	_, exists := labels[label]
	if exists {
		errorCodes = append(
			errorCodes,
			fmt.Sprintf("Line %d: Redefinition of label %s.", lex.line, label),
		)
	} else {
		labels[label] = labelInfo{address: address, line: lex.line}
	}
}

func setOrg(addr uint16) {
	setAddress(addr)
}

func addOpcodeLabel(opcode string, label string) {
	ops = append(
		ops,
		opInfo{op: opcode, address: address, label: label, line: lex.line},
	)
	setAddress(address + 2)
}

func addOpcodeConstant(opcode string, constant uint16) {
	ops = append(
		ops,
		opInfo{op: opcode, address: address, num: constant, line: lex.line},
	)
	setAddress(address + 2)
}

func addWords(newWords []uint16) {
	words = append(
		words,
		wordInfo{words: newWords, address: address, line: lex.line},
	)
	setAddress(address + uint16(len(newWords)))
}

func setAddress(addr uint16) {
	if addr > maxAddress {
		maxAddress = addr
	}
	address = addr
}
