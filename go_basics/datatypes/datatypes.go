package main

import (
	"fmt"
	"math/cmplx"
)

func getSignedIntegerExamples() {
	const defaultInteger int = 1

	// max: 127
	var eightBitInteger int8 = 127
	// max: 32767
	var sixteenBitInteger int16 = 32767

	// max: 2147483647
	var thirtyTwoBitInteger int32 = 2147483647

	// max: 9223372036854775807
	var sixtyFourBitInteger int64 = 9223372036854775807

	fmt.Printf(
		"Basic integer = %d, 8 bit = %d, 16bit = %d, 32 bit = %d, 64 bit = %d \n",
		defaultInteger, eightBitInteger, sixteenBitInteger, thirtyTwoBitInteger, sixtyFourBitInteger,
	)
}

func getUnsignedIntegerExamples() {
	var defaultInteger uint = 1

	var eightBitInteger uint8 = 255

	var sixteenBitInteger uint16 = 65535

	var thirtyTwoBitInteger uint32 = 2147483647

	var sixtyFourBitInteger uint64 = 9223372036854775807

	fmt.Printf(
		"Basic integer unsigned = %d,  8 bit unsigned = %d, 16bit unsigned = %d, 32 bit unsigned = %d, 64 bit unsigned = %d \n",
		defaultInteger, eightBitInteger, sixteenBitInteger, thirtyTwoBitInteger, sixtyFourBitInteger,
	)
}

func getSomeSpecialTypes() {

	var simpleByte byte = 1
	var simpleRune rune = 123456

	fmt.Printf("Simple byte = %d, Simple rune = %d \n", simpleByte, simpleRune)
}

func getFloatingNumbers() {

	var float32Bit float32 = 11231231231232133123213123123131232131.1001
	var float64bit float64 = 123456123213123213123123123123123213123213123.88138218318238123213

	fmt.Printf("float 32 = %f, float 64 = %f \n", float32Bit, float64bit)
}

func getComplexNumbers() {
	var complex64bit complex64 = complex64(cmplx.Sqrt(5 + 12i))
	var complex128bit complex128 = cmplx.Sqrt(12 + 24i)

	fmt.Printf("Complex 64 = %v, Complex 128 = %v \n", complex64bit, complex128bit)
}

func main() {

	// signed integers
	getSignedIntegerExamples()

	// unsigned integers
	getUnsignedIntegerExamples()

	// rune and bytes

	getSomeSpecialTypes()

	// floating numbers

	getFloatingNumbers()

	// complex numbers

	getComplexNumbers()

}
