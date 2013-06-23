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

/*
 vmtest runs a drones VM binary and provides it with a primitive terminal.

 Console I/O is conducted through buses 0-4.  Setting bus 1 to a non-zero value
 will print the low-order byte of bus 0 to the terminal, after which it will be
 set back to zero.

 Console input is provided one byte at a time in the low-order byte of bus 2,
 with bus 3 used as a flag to notify the VM that a new byte is available.  The
 VM should set it to zero to notify the console that it's ready for more input.

 Setting bus 4 to a non-zero value will terminate the program.
*/
package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"github.com/bieber/drones/vm"
	"os"
	"runtime"
)

func main() {
	verbose := flag.Bool(
		"v",
		false,
		"Display detailed VM state at each clock cycle.",
	)
	memSize := flag.Uint(
		"m",
		0xffff,
		"Memory size in 16-bit words.",
	)
	flag.Parse()

	if len(flag.Args()) != 1 {
		die(errors.New("No input file."))
	}
	file, err := os.Open(flag.Args()[0])
	if err != nil {
		die(err)
	}
	defer file.Close()
	stdin := bufio.NewReader(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	v := vm.New(uint16(*memSize), 5)
	err = v.LoadBinary(file)
	if err != nil {
		die(err)
	}

	bytesIn := make(chan byte)
	go readBytes(stdin, bytesIn)
	for {
		if *verbose {
			fmt.Println(v.Debug())
		}
		buf := make([]byte, 2)
		if v.Buses[3] == 0 {
			select {
			case c := (<-bytesIn):
				buf[0] = c
				v.Buses[2] = binary.LittleEndian.Uint16(buf)
				v.Buses[3] = 1
			default:
			}
		}
		if v.Buses[1] != 0 {
			binary.LittleEndian.PutUint16(buf, v.Buses[0])
			stdout.WriteByte(buf[0])
			stdout.Flush()
			v.Buses[1] = 0
		}
		if v.Buses[4] != 0 {
			break
		}
		v.Clock()
		runtime.Gosched()
	}
}

func readBytes(in *bufio.Reader, out chan byte) {
	for {
		c, err := in.ReadByte()
		if err == nil {
			out <- c
		}
	}
}

func die(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}
