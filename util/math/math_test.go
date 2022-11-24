package math

import (
	"testing"
)

func TestBytes(t *testing.T) {
	i := 1534814456789999998

	bs, err := ToBytes(i)
	if err != nil {
		t.Fatal("failed to ToBytes: ", err)
	}

	var j int
	err = ParseFromBytes(bs, &j)
	if err != nil {
		t.Fatal("failed to ParseFromBytes: ", err)
	}

	if i != j {
		t.Fatal("the values are different")
	}

	i = 1234

	bs, err = ToBytes(i)
	if err != nil {
		t.Fatal("failed to ToBytes: ", err)
	}

	err = ParseFromBytes(bs, &j)
	if err != nil {
		t.Fatal("failed to ParseFromBytes: ", err)
	}

	if i != j {
		t.Fatal("the values are different")
	}
}
