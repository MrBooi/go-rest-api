package main

import "fmt"

// Run - is going to be executed when you run the program
func Run() error {
	fmt.Println("Starting up our application...")
	return nil
}

func main() {
	fmt.Println("Hello, world!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
