package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var myIntVar int
	myIntVar = -12
	fmt.Println("My variable is: ", myIntVar)

	var myUnitVar uint
	myUnitVar = 12
	fmt.Println("my variable is :", myUnitVar)

	var myStringVar string
	myStringVar = "My string variable"
	fmt.Println("My variable is: ", myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println("My variable is: ", myBoolVar)

	fmt.Println("My variable address is: ", &myStringVar)

	myIntVar2 := 12
	fmt.Println("My variable is : ", myIntVar2)

	myStringVar2 := "My string variable with :="
	fmt.Println(myStringVar2)

	const myIntConst int = 12
	fmt.Println("My const is: ", myIntConst)

	const myStringConst string = "aaa"
	fmt.Println("My const is: ", myStringConst)

	fmt.Println()

	var my8BitsIntVar int8
	fmt.Printf("Int default value: %d\n", my8BitsIntVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8BitsIntVar, unsafe.Sizeof(my8BitsIntVar), unsafe.Sizeof(my8BitsIntVar)*8)

	var my16BitsIntVar int16
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my16BitsIntVar, unsafe.Sizeof(my16BitsIntVar), unsafe.Sizeof(my16BitsIntVar)*8)

	var my32BitsIntVar int32
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my32BitsIntVar, unsafe.Sizeof(my32BitsIntVar), unsafe.Sizeof(my32BitsIntVar)*8)

	var my64BitsIntVar int64
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my64BitsIntVar, unsafe.Sizeof(my64BitsIntVar), unsafe.Sizeof(my64BitsIntVar)*8)

	fmt.Println()

	var my8BitsUintVar uint8
	fmt.Printf("Uint default value: %d\n", my8BitsUintVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8BitsUintVar, unsafe.Sizeof(my8BitsUintVar), unsafe.Sizeof(my8BitsUintVar)*8)

	var my16BitsUintVar uint16
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my16BitsUintVar, unsafe.Sizeof(my16BitsUintVar), unsafe.Sizeof(my16BitsUintVar)*8)

	var my32BitsUintVar uint32
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my32BitsUintVar, unsafe.Sizeof(my32BitsUintVar), unsafe.Sizeof(my32BitsUintVar)*8)

	var my64BitsUintVar uint64
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my64BitsUintVar, unsafe.Sizeof(my64BitsUintVar), unsafe.Sizeof(my64BitsUintVar)*8)

	fmt.Println()

	var myFloat32Var float32
	fmt.Printf("Float default value: %f\n", myFloat32Var)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

	var myFloat64Var float64
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)

	fmt.Println()

	var myStringVar3 string
	fmt.Printf("string default value: %s\n", myStringVar3)

	myStringVar5 := `My string variable in golang
	with multiple
	line`

	fmt.Printf("string default value: %s\n", myStringVar5)

	{
		fmt.Println()
		floatVar := 33.11
		fmt.Printf("type: %T, value: %f\n", floatVar, floatVar)
		floatStrVar := fmt.Sprintf("%.2f", floatVar)
		fmt.Printf("type: %T, value: %s\n", floatStrVar, floatStrVar)

		intVar := 22
		fmt.Printf("type: %T, value: %d\n", intVar, intVar)
		intStrVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("type: %T, value: %s\n", intStrVar, intStrVar)

		intVal1, err := strconv.ParseInt("1333", 0, 64)
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", intVal1, intVal1)

		intVal2, err := strconv.ParseInt("aa122", 0, 64)
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", intVal2, intVal2)

		floatVar1, _ := strconv.ParseFloat("-11.2", 64)
		fmt.Printf("type: %T, value: %f\n", floatVar1, floatVar1)
	}

	fmt.Println()
}
