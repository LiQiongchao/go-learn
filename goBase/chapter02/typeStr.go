package main

import (
	"fmt"
)

func sayHello(s string) string {
	return "Hello " + s
}

func main() {
	fmt.Printf(sayHello("George"))
}
