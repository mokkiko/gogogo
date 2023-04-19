package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// boolean => deafault = false
	var firstBoolean bool
	fmt.Println((firstBoolean))
	aBoolean, bBoolean := true, true
	fmt.Println("AND = ", aBoolean && bBoolean)
	fmt.Println(" OR = ", aBoolean || bBoolean)
	fmt.Println("NOT = ", !aBoolean)

	// Numerics. Integer
	// int8, int16, int32, int64, int - число после  - биты, в 1 бите 8 байтов
	// uint8, uint16, uint32, uint64, int

	var a int = 32
	b := 92
	fmt.Println(a + b)
	// Узнать сколько байт занимает перменная типа int
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a))

}
