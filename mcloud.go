package uuid

import (
	"github.com/globalsign/mgo/bson"
)

// SetBSON allows the struct to be correctly retrieved from mongo
func (u *UUID) SetBSON(raw bson.Raw) error {
	var doc bson.Binary
	raw.Unmarshal(&doc)
	copy(u[:], doc.Data)
	return nil
}

// GetBSON allows the struct to be correctly stored as mongo binary
func (u *UUID) GetBSON() (interface{}, error) {
	bytes, err := u.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return bson.Binary{
		Kind: 4,
		Data: bytes,
	}, nil
}

func (u *UUID) UnmarshalParam(s string) error {
	id, err := Parse(s)
	*u = UUID(id)
	return err
}
