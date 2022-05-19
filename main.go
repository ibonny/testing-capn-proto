package main

import (
	"os"
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

	// Write the message to stdout.
	if err = capnp.NewEncoder(os.Stdout).Encode(msg); err != nil {
		panic(err)
	}
}
