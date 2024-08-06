//go:build ignore

package main

import (
	"fmt"
)

func main() {
	buffer := [2]byte{}

	buffer[0] = 105
	buffer[1] = 115
	str1 := "vaibhav"
	str2 := "kesarkar"
	// fmt.Println(max(len(str1),len(str2)))

	fmt.Println("c1:", str1[2]|str2[2])
	// for i,c := range(max(str1,))

	fmt.Println(buffer[0] | buffer[1])
	fmt.Println(buffer[0] & buffer[1])
	fmt.Println(^buffer[0])
	fmt.Println(buffer[0] >> 1)
	fmt.Println(buffer[1] << 2)

	// fmt.Println(len(buffer),cap(buffer))
	fmt.Printf("%T", buffer)
}
