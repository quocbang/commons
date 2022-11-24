package math

import (
	"bytes"
	"encoding/binary"
)

// int極值
const (
	MaxInt = int((^uint(0)) >> 1)
	MinInt = ^(int((^uint(0)) >> 1))
)

// ToBytes returns the bytes representation of data.
func ToBytes(data interface{}) ([]byte, error) {
	if i, ok := data.(int); ok {
		if uint(i)>>32 > 0 {
			data = int64(i)
		} else {
			data = int32(i)
		}
	}

	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ParseFromBytes parses the bytes into p.
func ParseFromBytes(data []byte, p interface{}) error {
	buf := bytes.NewBuffer(data)
	if pp, ok := p.(*int); ok {
		if buf.Len() == 4 {
			var i int32
			err := binary.Read(buf, binary.LittleEndian, &i)
			if err != nil {
				return err
			}
			*pp = int(i)
		} else {
			var i int64
			err := binary.Read(buf, binary.LittleEndian, &i)
			if err != nil {
				return err
			}
			*pp = int(i)
		}
		return nil
	}
	return binary.Read(buf, binary.LittleEndian, p)
}

// MustToBytes returns the bytes representation of data.
// It panics if gets an error.
func MustToBytes(data interface{}) []byte {
	bs, err := ToBytes(data)
	if err != nil {
		panic(err)
	}
	return bs
}

// MustParseFromBytes parses the bytes into p.
// It panics if gets an error.
func MustParseFromBytes(data []byte, p interface{}) {
	err := ParseFromBytes(data, p)
	if err != nil {
		panic(err)
	}
}
