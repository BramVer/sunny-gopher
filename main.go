package main

// Import declares lib packages
import (
	"fmt"       // A package in standard Go libr
	"io/ioutil" // Implements some io util functions
	"os"        // OS functions
	"strconv"   // String conversions
)

func main() {
	// Call pkg name before calling func
	fmt.Println("HENLO FRENDO!!")

	// Call another func within THIS pkg
	beyondHenlo()
}

// Funcs have no params in parantheses!!!
func beyondHenlo() {

	// Long var declaration
	var x int
	x = 3

	// Short var declaration
	y := 3

	sum, prod := learnMultiple(x, y)

	fmt.Println("Sum: ", sum, " -- Prod: ", prod)
	learnTypes()
}
