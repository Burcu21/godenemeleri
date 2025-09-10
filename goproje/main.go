package main

import (
	"fmt"
	"projegodeneme/c1"
	"projegodeneme/d1"
)

func main() {

	sum := c1.Deneme(3, 5)
	fmt.Println("Sum:", sum)

	multiplication := c1.Carpma(4, 6)
	fmt.Println("Multiplication:", multiplication)

	division := c1.Bolme(20, 4)
	fmt.Println("Division:", division)

	sub := c1.Cikarma(10, 3)
	fmt.Println("Subtraction:", sub)

	name := d1.Name("Burcu")
	fmt.Println("Name:", name)

}
