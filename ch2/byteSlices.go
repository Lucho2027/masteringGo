package main

import "fmt"

func main() {
	b := make([]byte, 12)

	fmt.Println("Byte slice:", b)

	b = []byte("Byteslice $")
	fmt.Println("Byte slice:", b)

	fmt.Printf("Byte slice as text : %s\n", b)
	fmt.Println("Byte slie as text:", string(b))

	fmt.Println("LEngth of b:", len(b))
}
