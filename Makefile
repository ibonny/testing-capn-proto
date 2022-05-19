all: test

schemas/books.capnp.go: schemas/books.capnp
	capnp compile -I $$GOPATH/src/capnproto.org/go/capnp/std -ogo schemas/books.capnp

test: main.go schemas/books.capnp.go
	go run main.go