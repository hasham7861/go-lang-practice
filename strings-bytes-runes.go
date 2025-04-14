// Note strings can hold utf-8 and unicode bytes internally but string litterals are just arbitary text with utf-8 characters and non with escape and unicode values.
// But the string literal itself is always in utf-8 readable format but the container itself can hold utf-8 string and unicode
//
// strings are stored as utf-8 and it relies on unicode to convert it to utf-8 byte or bytes depending on the codepoint aka rune. All the ascii values are one code point and also one hexdecimal utf-8 single byte, but not same for character that dont live in utf-8 but are unicode where by it has multiple bytes
package main

import "fmt"

// func main() {
// 	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

// 	fmt.Println("Println:")
// 	// attempts to print out bytes, by deault it atempts to output to utf-8 acsii where by you will see weird symbols show up as it is not understood
// 	fmt.Println(sample)

// 	fmt.Println("Byte loop:")
// 	for i := 0; i < len(sample); i++ {
// 		// % in Printf prints out each char as a hexdecimal byte value
// 		fmt.Printf("%x ", sample[i])
// 	}
// 	fmt.Printf("\n")

// 	fmt.Println("Printf with %x:")
// 	// you can print entire string hexadecimal byte values as they are
// 	fmt.Printf("%x\n", sample)

// 	fmt.Println("Printf with % x:")
// 	// you can put space between % x to put spaces between the utf-8 byte encodings
// 	fmt.Printf("% x\n", sample)

// 	fmt.Println("Printf with %q:")
// 	// includes the escap characters such as \x for byte notation and new line \n notation
// 	fmt.Printf("%q\n", sample)

// 	fmt.Println("Printf with %+q:")
// 	// with plus notation with q and percent then this includes all characters that are utf-8 encoded and non utf-8 encoded bytes
// 	fmt.Printf("%+q\n", sample)
// }

// place of interest symbole printed
// output in unicode with %+q if not utf-8 and then %x for byte at a time
// plain string: ⌘
// quoted string: "\u2318"
// hex bytes: e2 8c 98
func main() {
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")
}
