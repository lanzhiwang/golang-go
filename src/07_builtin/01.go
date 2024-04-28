package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var var1 uint8
	var var2 uint16
	var var3 uint32
	var var4 uint64

	var var5 int8
	var var6 int16
	var var7 int32
	var var8 int64

	var var9 float32
	var var10 float64

	var var11 string

	var var12 int

	var var13 uint

	var var14 byte // type byte = uint8

	var var15 rune // type rune = int32

	// var var16 IntegerType // type IntegerType int
	// var var17 Type        // type Type int
	// var var18 Type1       // type Type1 int

	// var var19 FloatType // type FloatType float32

	fmt.Printf("var1 的数据类型是 %T, 占的字节数是 %d\n", var1, unsafe.Sizeof(var1))
	fmt.Printf("var2 的数据类型是 %T, 占的字节数是 %d\n", var2, unsafe.Sizeof(var2))
	fmt.Printf("var3 的数据类型是 %T, 占的字节数是 %d\n", var3, unsafe.Sizeof(var3))
	fmt.Printf("var4 的数据类型是 %T, 占的字节数是 %d\n", var4, unsafe.Sizeof(var4))
	fmt.Printf("var5 的数据类型是 %T, 占的字节数是 %d\n", var5, unsafe.Sizeof(var5))
	fmt.Printf("var6 的数据类型是 %T, 占的字节数是 %d\n", var6, unsafe.Sizeof(var6))
	fmt.Printf("var7 的数据类型是 %T, 占的字节数是 %d\n", var7, unsafe.Sizeof(var7))
	fmt.Printf("var8 的数据类型是 %T, 占的字节数是 %d\n", var8, unsafe.Sizeof(var8))
	fmt.Printf("var9 的数据类型是 %T, 占的字节数是 %d\n", var9, unsafe.Sizeof(var9))
	fmt.Printf("var10 的数据类型是 %T, 占的字节数是 %d\n", var10, unsafe.Sizeof(var10))
	fmt.Printf("var11 的数据类型是 %T, 占的字节数是 %d\n", var11, unsafe.Sizeof(var11))
	fmt.Printf("var12 的数据类型是 %T, 占的字节数是 %d\n", var12, unsafe.Sizeof(var12))
	fmt.Printf("var13 的数据类型是 %T, 占的字节数是 %d\n", var13, unsafe.Sizeof(var13))
	fmt.Printf("var14 的数据类型是 %T, 占的字节数是 %d\n", var14, unsafe.Sizeof(var14))
	fmt.Printf("var15 的数据类型是 %T, 占的字节数是 %d\n", var15, unsafe.Sizeof(var15))
	// fmt.Printf("var16 的数据类型是 %T, 占的字节数是 %d\n", var16, unsafe.Sizeof(var16))
	// fmt.Printf("var17 的数据类型是 %T, 占的字节数是 %d\n", var17, unsafe.Sizeof(var17))
	// fmt.Printf("var18 的数据类型是 %T, 占的字节数是 %d\n", var18, unsafe.Sizeof(var18))
	// fmt.Printf("var19 的数据类型是 %T, 占的字节数是 %d\n", var19, unsafe.Sizeof(var19))

	fmt.Println("int8:", math.MinInt8, "~", math.MaxInt8)
	fmt.Println("int16:", math.MinInt16, "~", math.MaxInt16)
	fmt.Println("int32:", math.MinInt32, "~", math.MaxInt32)
	fmt.Println("int64:", math.MinInt64, "~", math.MaxInt64)

	fmt.Println("uint8:", "0", "~", math.MaxUint8)
	fmt.Println("uint16:", "0", "~", math.MaxUint16)
	fmt.Println("uint32:", "0", "~", math.MaxUint32)
	// fmt.Println("uint64:", "0", "~", math.MaxUint64)

	fmt.Println("int:", math.MinInt, "~", math.MaxInt)

	// fmt.Println("uint:", "0", "~", math.MaxUint)

	v1 := 'a' // v1 的数据类型是 int32, 占的字节数是 4  type rune = int32
	v2 := "a" // v2 的数据类型是 string, 占的字节数是 16
	// v3 := 'ab' // more than one character in rune literal
	v4 := "ab" // v4 的数据类型是 string, 占的字节数是 16
	v5 := '中'  // v5 的数据类型是 int32, 占的字节数是 4  type rune = int32
	v6 := "中"  // v6 的数据类型是 string, 占的字节数是 16
	// v7 := '中文' // more than one character in rune literal
	v8 := "中文" // v8 的数据类型是 string, 占的字节数是 16
	fmt.Printf("v1 的数据类型是 %T, 占的字节数是 %d\n", v1, unsafe.Sizeof(v1))
	fmt.Printf("v2 的数据类型是 %T, 占的字节数是 %d\n", v2, unsafe.Sizeof(v2))
	fmt.Printf("v4 的数据类型是 %T, 占的字节数是 %d\n", v4, unsafe.Sizeof(v4))
	fmt.Printf("v5 的数据类型是 %T, 占的字节数是 %d\n", v5, unsafe.Sizeof(v5))
	fmt.Printf("v6 的数据类型是 %T, 占的字节数是 %d\n", v6, unsafe.Sizeof(v6))
	fmt.Printf("v8 的数据类型是 %T, 占的字节数是 %d\n", v8, unsafe.Sizeof(v8))

	var b1 byte = 0   // b1 的数据类型是 uint8, 占的字节数是 1
	var b2 byte = 255 // b2 的数据类型是 uint8, 占的字节数是 1
	// var b3 byte = -1 // cannot use -1 (untyped int constant) as byte value in variable declaration (overflows)
	// var b4 byte = 256 // cannot use 256 (untyped int constant) as byte value in variable declaration (overflows)
	var b5 byte = 'a' // b5 的数据类型是 uint8, 占的字节数是 1
	// var b6 byte = "a" // cannot use "a" (untyped string constant) as byte value in variable declaration
	b7 := []byte("abcde")
	b8 := []byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}
	fmt.Printf("b1 的数据类型是 %T, 占的字节数是 %d\n", b1, unsafe.Sizeof(b1))
	fmt.Printf("b2 的数据类型是 %T, 占的字节数是 %d\n", b2, unsafe.Sizeof(b2))
	fmt.Printf("b5 的数据类型是 %T, 占的字节数是 %d\n", b5, unsafe.Sizeof(b5))
	fmt.Printf("%v %q\n", b7, b7) // [97 98 99 100 101] "abcde"
	fmt.Printf("%v %q\n", b8, b8) // [104 101 108 108 111 32 119 111 114 108 100] "hello world"

}
