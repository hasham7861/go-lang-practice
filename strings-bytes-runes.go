package main

import "fmt"

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	// attempts to print out bytes, by deault it atempts to output to utf-8 acsii where by you will see weird symbols show up as it is not understood
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		// % in Printf prints out each char as a hexdecimal byte value
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	// you can print entire string hexadecimal byte values as they are
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	// you can put space between % x to put spaces between the utf-8 byte encodings
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	// includes the escap characters such as \x for byte notation and new line \n notation
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	// with plus notation with q and percent then this includes all characters that are utf-8 encoded and non utf-8 encoded bytes
	fmt.Printf("%+q\n", sample)
}
