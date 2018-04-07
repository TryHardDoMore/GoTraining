// Solution to part 1 of the Whispering Gophers code lab.
// This program reads from standard input and writes JSON-encoded messages to
// standard output. For example, this input line:
//	Hello!
// Produces this output:
//	{"Body":"Hello!"}
//
package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type Message struct {
	Body string
}

func main() {
	// TODO: Create a new bufio.Scanner reading from the standard input.
	scanner := bufio.NewScanner(os.Stdin)
	// TODO: Create a new json.Encoder writing into the standard output.
	encoder := json.NewEncoder(os.Stdout)
	/* TODO: Iterate over every line in the scanner */
	for scanner.Scan() {
		// TODO: Create a new message with the read text.
		mes := Message{Body: scanner.Text()}
		// TODO: Encode the message, and check for errors!
		encoder.Encode(mes)
	}
	// TODO: Check for a scan error.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
