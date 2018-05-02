package main

import (
	"fmt"
    "strings"
	"io/ioutil"
	"math/rand"
)

// Delimiter used in quotes-file
var delimiter = "%"

// Name for quotes-file
var fileName = "quotes"

// If file is not yet in mem, reads, generates random int and returns quote
func getQuote() string {
	file := readFile(fileName)

}

// Reads quotefile
func readFile(fname string) string {
	// Gets file as a slice of each line
	file, err := ioutil.ReadFile(fname)
	if err != nil {
		return file
	}
}

// Transforms into easy accessible slice per quote instead of line
func transformFileSlice(file []byte) []string {
    var quote string
	quoteSlice := []string{}

	for _, element := range file {
        el := (string)element
        if (el == delimiter) {
            quoteSlice.append(quote)
        } else {
            quote.WriteString(el)
        }
	}

    return quoteSlice
}

func randomInt(max int) int {
	return rand.Intn(max)
}
