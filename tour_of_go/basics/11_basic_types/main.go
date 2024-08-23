package main

import "fmt"

func main() {
	var boolVal bool = true
	fmt.Printf("boolVal = '%v'\n", boolVal)

	var strVal string = "Hello, world"
	fmt.Printf("strVal = '%s'\n", strVal)

	// for int or intX types, the values are signed,
	// which means 1 bit is used to determine the sign
	// of the value, positive, or negative!
	var intVal int = 123
	var int8Val int8 = 127                   // cannot be more than (2^7 - 1)
	var int16Val int16 = 32767               // cannot be more than (2^15 - 1)
	var int32Val int32 = 2147483647          // cannot be more than (2^31 - 1)
	var int64Val int64 = 9223372036854775807 // cannot be more than (2^63 - 1)
	fmt.Printf("intVal = %v\n", intVal)
	fmt.Printf("int8Val = %v\n", int8Val)
	fmt.Printf("int16Val = %v\n", int16Val)
	fmt.Printf("int32Val = %v\n", int32Val)
	fmt.Printf("int64Val = %v\n", int64Val)

	// uints follow the same rules as ints,
	// except we can use 1 more bit,
	// which means the magnitude of the allowed values
	// increases by a factor of 2.
	var uintVal uint = 123
	var uint8Val uint8 = 255                    // cannot be more than (2^8 - 1)
	var uint16Val uint16 = 65535                // cannot be more than (2^16 - 1)
	var uint32Val uint32 = 4294967295           // cannot be more than (2^32 - 1)
	var uint64Val uint64 = 18446744073709551615 // cannot be more than (2^64 - 1)
	fmt.Printf("uintVal = %v\n", uintVal)
	fmt.Printf("uint8Val = %v\n", uint8Val)
	fmt.Printf("uint16Val = %v\n", uint16Val)
	fmt.Printf("uint32Val = %v\n", uint32Val)
	fmt.Printf("uint64Val = %v\n", uint64Val)

	fmt.Println("Skipping using 'uintptr'...")
	// var uintptrVal uintptr = 234 // im not gonna use this.
	// no systems-level programming here yet. im a beginner

	var byteVal byte = 255 // 'byte' is an alias for uint8
	fmt.Printf("byteVal = %v\n", byteVal)

	var runeVal rune = 2147483647 // 'rune' is an alias for int32
	// runes represent a Unicode code point.
	fmt.Printf("runeVal = %v\n", runeVal)

	var float32Val float32 = 123.45 // Range is -3.4e+38 to 3.4e+38.
	var float64Val float64 = 543.21 // Range is -1.7e+308 to +1.7e+308.
	fmt.Printf("float32Val = %v\n", float32Val)
	fmt.Printf("float64Val = %v\n", float64Val)

	var complex64Val complex64 = 6 + 2i
	var complex128Val complex128 = 1 - 9i
	fmt.Printf("complex64Val = %v\n", complex64Val)
	fmt.Printf("complex128Val = %v\n", complex128Val)

}
