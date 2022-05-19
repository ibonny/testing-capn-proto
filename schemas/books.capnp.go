// Code generated by capnpc-go. DO NOT EDIT.

package schemas

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
)

type Book struct{ capnp.Struct }

// Book_TypeID is the unique identifier for the type Book.
const Book_TypeID = 0x8100cc88d7d4d47c

func NewBook(s *capnp.Segment) (Book, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Book{st}, err
}

func NewRootBook(s *capnp.Segment) (Book, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Book{st}, err
}

func ReadRootBook(msg *capnp.Message) (Book, error) {
	root, err := msg.Root()
	return Book{root.Struct()}, err
}

func (s Book) String() string {
	str, _ := text.Marshal(0x8100cc88d7d4d47c, s.Struct)
	return str
}

func (s Book) Title() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Book) HasTitle() bool {
	return s.Struct.HasPtr(0)
}

func (s Book) TitleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Book) SetTitle(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Book) PageCount() int32 {
	return int32(s.Struct.Uint32(0))
}

func (s Book) SetPageCount(v int32) {
	s.Struct.SetUint32(0, uint32(v))
}

// Book_List is a list of Book.
type Book_List = capnp.StructList[Book]

// NewBook creates a new list of Book.
func NewBook_List(s *capnp.Segment, sz int32) (Book_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return capnp.StructList[Book]{List: l}, err
}

// Book_Future is a wrapper for a Book promised by a client call.
type Book_Future struct{ *capnp.Future }

func (p Book_Future) Struct() (Book, error) {
	s, err := p.Future.Struct()
	return Book{s}, err
}

const schema_85d3acc39d94e0f8 = "x\xda\x12Hr`1\xe4\xdd\xcf\xc8\xc0\x14(\xc2\xca" +
	"\xb6\xbf\xe6\xca\x95\xeb\x1dg\x1a\x03E\x18\x19\xff\xffx" +
	"0e\xee\xe15\x97[\x19X\x19\xd9\x19\x18\x04\x8f\x9e" +
	"\x12\xbc\x08\xa2\xcf\x963\xe8\xfe/N\xceH\xcdM," +
	"\xd6gJ\xca\xcf\xcf.\xd6KN,\xc8+\xb0r\xca" +
	"\xcf\xcff`\x08`d\x0c\xe4`fa``ad" +
	"`\x10\xd44b`\x08Taf\x0c4`bdd" +
	"\x14a\x04\x89\xe9\x0610\x04\xea03\x06Z01" +
	"\xca\x97d\x96\xe4\xa42\xf2001\xf200\xfe/" +
	"HLOu\xce/\xcdc`,ada`bd" +
	"a`\x04\x04\x00\x00\xff\xffO$(!"

func init() {
	schemas.Register(schema_85d3acc39d94e0f8,
		0x8100cc88d7d4d47c)
}