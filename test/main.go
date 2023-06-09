package main

import "fmt"

type Demo struct {
	name string
}

func getDemo() *Demo {
	return &Demo{
		name: "mika",
	}
}

func printName(demo *Demo) {
	fmt.Println(demo.name)

}

func main() {
	demo := getDemo()
	printName(demo)
}
