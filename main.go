package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Printf("Let's begin!\n")

	var name string

	var run = flag.String("r", "none", "Name of the function you want to run.")
	flag.StringVar(&name, "n", "none", "Some name to say hello!")

	flag.Parse()

	fmt.Println("Running: ", *run)

	switch *run {
	case "hello":
		hello(name)
	default:
		fmt.Println("Hi ..")

	}

}
func hello(name string) {
	fmt.Printf("Hello %s!\n", name)
}
