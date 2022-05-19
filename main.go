package main

import (
	"bytes"
	"fmt"
	"testing_capn_proto/schemas"

	"capnproto.org/go/capnp/v3"
)

func main() {
	// Make a brand new empty message.  A Message allocates Cap'n Proto structs.
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	// Create a new Book struct.  Every message must have a root struct.
	book, err := schemas.NewRootBook(seg)
	if err != nil {
		panic(err)
	}

	book.SetTitle("War and Peace")
	book.SetPageCount(1440)

	var buff bytes.Buffer

	// Write the message to stdout.
	if err = capnp.NewEncoder(&buff).Encode(msg); err != nil {
		panic(err)
	}

	var msg2 *capnp.Message

	// Read the message from stdin.
	if msg2, err = capnp.NewDecoder().Decode(); err != nil {
		panic(err)
	}

	// Extract the root struct from the message.
	book, err := books.ReadRootBook(msg)
	if err != nil {
		panic(err)
	}

	// Access fields from the struct.
	title, err := book.Title()
	if err != nil {
		panic(err)
	}
	pageCount := book.PageCount()
	fmt.Printf("%q has %d pages\n", title, pageCount)
}
