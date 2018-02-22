package aSimplePackage

import "fmt"

const Pi = "3.14159"
const version = "1.1"

func init() {
	fmt.Println("The init function of aSimplePackage")
}

func Add(x, y int) int {
	return x + y
}

func Println(x int) {
	fmt.Println(x)
}

func Version() {
	fmt.Println("The version of the package is", version)
}
