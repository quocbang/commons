package kendaql

import (
	"reflect"
	"testing"
	"time"
)

func TestBool(t *testing.T) {
	data := true

	bs, err := ToValue(data)
	if err != nil {
		t.Fatal(err)
	}

	val, err := ValueToBool(bs)
	if err != nil {
		t.Fatal(err)
	}

	if val != data {
		t.Fatalf("got %v, expected %v", val, data)
	}
}
func TestBytes(t *testing.T) {
	data := []byte("test")

	bs, err := ToValue(data)
	if err != nil {
		t.Fatal(err)
	}

	val, err := ValueToBytes(bs)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(val, data) {
		t.Fatalf("got %v, expected %v", val, data)
	}
}

func TestTime(t *testing.T) {
	data, err := time.Parse("2006-01-02 15:04:05", "2019-12-31 14:03:59")

	bs, err := ToValue(data)
	if err != nil {
		t.Fatal(err)
	}

	val, err := ValueToTime(bs)
	if err != nil {
		t.Fatal(err)
	}

	if val != data {
		t.Fatalf("got %v, expected %v", val, data)
	}
}
