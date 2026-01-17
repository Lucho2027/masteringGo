package main

import "fmt"

type aStructure struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 12.123
	fmt.Println("Memory address of f:", &f)

	// Pointer to f
	fP := &f
	fmt.Println("Memory address of f:", &f)
	fmt.Println("Value of f:", *fP)

	// the value of changes
	processPointer(fP)
	fmt.Printf("Value of f: %.2f\n", f)
	// The vaue of f does not changes
	x := returnPointer(f)
	fmt.Printf("Value of x: %.2f\n", *x)
	// The value of f does not changes
	xx := bothPointers(fP)
	fmt.Printf("Value of xx: %.2f\n", *xx)

	// Check for empty aStructure
	var k *aStructure
	fmt.Println("k here should be nil", k)

	// Therefore you are allowed to do this:

	if k == nil {
		k = new(aStructure)
	}
	fmt.Printf("%+v\n", k)
	if k != nil {
		fmt.Println("k is not nill")
	}
}
