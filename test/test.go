package main

import (
	"fmt"
	"github.com/markedhero/configurit"
	"os"
)

func main() {
	fmt.Println(configurit.Version())

	c, err := configurit.CreateFrom("./config.conf")
	if err != nil {
		fmt.Printf("There was a problem: %s\n", err)
		os.Exit(1)
	}

	a := c.GetSectionNames()
	b, err := c.GetSection(a[2])
	if err != nil {
		fmt.Printf("There was a problem: %s\n", err)
		os.Exit(1)
	}
	b.Print()
	fmt.Println()

	test, err := b.GetBool("option")
	if err != nil {
		fmt.Printf("There was a problem: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(test)
}
