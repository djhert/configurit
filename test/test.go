package main

import (
	"fmt"
	"github.com/hlfstr/configurit"
	"os"
)

func main() {
	fmt.Println(configurit.Version())

	c, err := configurit.Open("./config.conf")
	if err != nil {
		fmt.Printf("There was a problem: %s\n", err)
		os.Exit(1)
	}

	c.Print()
	fmt.Println(c.Config["Aeridya"]["UseSSL"])
}
