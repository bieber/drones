/*
 * Copyright 2013, Robert Bieber
 *
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
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"github.com/bieber/drones/asm"
	"io"
	"os"
)

func main() {
	outPath := flag.String("o", "-", "Output file path.  Use - for stdout.")
	flag.Parse()
	var inPath string
	if len(flag.Args()) == 0 {
		inPath = "-"
	} else if len(flag.Args()) == 1 {
		inPath = flag.Args()[0]
	} else {
		die(errors.New("Too many arguments"))
	}

	var fin *bufio.Reader
	var fout io.Writer
	if inPath == "-" {
		fin = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.Open(inPath)
		if err != nil {
			die(err)
		}
		fin = bufio.NewReader(file)
		defer file.Close()
	}
	if *outPath == "-" {
		fout = os.Stdout
	} else {
		file, err := os.Create(*outPath)
		if err != nil {
			die(err)
		}
		defer file.Close()
		fout = file
	}

	code, err := asm.Assemble(fin)
	if err != nil {
		die(err)
	}
	binary.Write(fout, binary.LittleEndian, uint16(len(code)))
	binary.Write(fout, binary.LittleEndian, code)
}

func die(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}
