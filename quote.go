package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Delimiter used in quotes-file
var delimiter = "%"

// Name for quotes-file
var fileName = "quotes"

// Keeps string as shared var to prevent rereading
var file string

// If file is not yet in mem, reads, generates random int and returns quote
func getQuote() string {
	if file == "" {
		file = readFile("assets/" + fileName)
	}
	quotes := transformQuoteSlice(file)

	// Determine random value based on total amount of quotes
	rand := randomInt(len(quotes))
	quote := quotes[rand-1]

	return quote
}

// Reads quotefile
func readFile(fname string) string {
	// Gets file as a slice of each line
	file, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	return string(file)
}

// Transforms into easy accessible slice per quote instead of line
func transformQuoteSlice(file string) []string {
	quoteSlice := strings.Split(file, delimiter)
	return quoteSlice
}

func randomInt(max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max)
}
