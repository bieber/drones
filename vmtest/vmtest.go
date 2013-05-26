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
