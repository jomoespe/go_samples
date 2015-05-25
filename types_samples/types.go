/* Go types examples
 * 
 * https://golang.org/ref/spec#Types
 */
package main

import (
	"fmt"
)

// Boolean types
type bol1  bool

// Numeric types
type num1  uint8       // the set of all unsigned  8-bit integers (0 to 255)
type num2  uint16      // the set of all unsigned 16-bit integers (0 to 65535)
type num3  uint32      // the set of all unsigned 32-bit integers (0 to 4294967295)
type num4  uint64      // the set of all unsigned 64-bit integers (0 to 18446744073709551615)

type num5  int8        // the set of all signed  8-bit integers (-128 to 127)
type num6  int16       // the set of all signed 16-bit integers (-32768 to 32767)
type num7  int32       // the set of all signed 32-bit integers (-2147483648 to 2147483647)
type num8  int64       // the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

type num9  float32     // the set of all IEEE-754 32-bit floating-point numbers
type num10 float64     // the set of all IEEE-754 64-bit floating-point numbers

type num11 complex64   // the set of all complex numbers with float32 real and imaginary parts
type num12 complex128  // the set of all complex numbers with float64 real and imaginary parts

type num13 byte        // alias for uint8
type num14 rune        // alias for int32

type num15 uint        // either 32 or 64 bits
type num16 int         // same size as uint
type num17 uintptr     // an unsigned integer large enough to store the uninterpreted bits of a pointer value

// String types
type str1 string 	   // string

// array types
type array1 [32]byte
type array2 [2][2][2]float32 // n-dimensional array

// structs
/*
struct {}					// an empty struct

struct {				  // a struct with 6 fields
	x, y int
	u    float32
	_    float32          //   padding
	A   *[]int
	F   func()
}
*/

func main() {
	fmt.Printf("character %s starts at byte position %s\n", "", "")
}