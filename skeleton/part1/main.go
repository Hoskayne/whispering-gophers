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
	//TODO: Create a new bufio.Scanner reading from the standard input.
	enc := json.NewEncoder(os.Stdout)
	//TODO: Create a new json.Encoder writing into the standard output.
	scanner := bufio.NewScanner(os.Stdin)
	var msgs []Message

	//TODO: Iterate over every line in the scanner
	for scanner.Scan() && (scanner.Text() != "q") {
		//TODO: Create a new message with the read text.
		var msg Message
		msg.Body = scanner.Text()
		msgs = append(msgs, msg)
	}
	//TODO: Check for a scan error.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//TODO: Encode the message, and check for errors!
	for _, s := range msgs {
		err := enc.Encode(s)
		if err != nil {
			log.Fatal(err)
		}
	}
}

/******SOLUTION:
func main() {
	s := bufio.NewScanner(os.Stdin)
	e := json.NewEncoder(os.Stdout)
	for s.Scan() {
		m := Message{Body: s.Text()}
		err := e.Encode(m)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
*/
