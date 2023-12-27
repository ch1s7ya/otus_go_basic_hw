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

func (b *BookSlice) MarshalSlice() ([]byte, error) {
	return proto.Marshal(b)
}

func (b *BookSlice) UnmarshalSlice(data []byte) error {
	return proto.Unmarshal(data, b)
}
