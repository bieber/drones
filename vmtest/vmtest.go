// vmtest runs test binaries through the VM and debugs along the way.
package main

import (
	"fmt"
	"github.com/bieber/drones/vm"
)

func main() {
	v := vm.New(20, 10)
	fmt.Println(v.Debug())
}