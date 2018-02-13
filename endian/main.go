package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int = 0x012345678
	var p unsafe.Pointer = unsafe.Pointer(&x)

	var bp *[4]byte = (*[4]byte)(p)

	var endian string = "unknown"

	if 0x01 == bp[0] {
		endian = "big"
	} else if (0x78 & 0xff) == (bp[0] & 0xff) {
		endian = "little"
	} else {
		var i uintptr = 0
		for i = 0; i < unsafe.Sizeof(x); i++ {
			fmt.Printf("byte[%d] = %d\n", i, bp[i])
		}
	}
	fmt.Printf("Machine has %s endian (int is %d bits)\n", endian, unsafe.Sizeof(x)*8)
}
