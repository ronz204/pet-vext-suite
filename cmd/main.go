package main

import "os"

func main() {
	if len(os.Args) < 2 {
		println("No command provided")
	} else {
		println("Command provided:", os.Args[1])
	}
}
