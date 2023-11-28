package protobuf

import (
	"google.golang.org/protobuf/proto"
)

func (r *Book) Marshal() ([]byte, error) {
	return proto.Marshal(r)
}

func (r *Book) Unmarshal(data []byte) error {
	return proto.Unmarshal(data, r)
}

func (b *Bookshelf) MarshalSlice() ([]byte, error) {
	return proto.Marshal(b)
}

func (b *Bookshelf) UnmarshalSlice(data []byte) error {
	return proto.Unmarshal(data, b)
}
