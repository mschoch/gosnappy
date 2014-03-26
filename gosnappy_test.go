package gosnappy

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEncodeDecodeRoundTrip(t *testing.T) {

	inputs := [][]byte{
		bytes.Repeat([]byte{'a'}, 128),
	}

	for _, input := range inputs {
		output, err := Encode(input)
		if err != nil {
			t.Fatal(err)
		}
		roundTrip, err := Decode(output)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(input, roundTrip) {
			t.Errorf("Mismatch - Input: %# x RoundTripped: %# x", input, roundTrip)
		}
	}
}

func TestDecodeEncodeRoundTrip(t *testing.T) {

	inputs := [][]byte{
		[]byte{0x80, 0x01, 0x00, 0x61, 0xfe, 0x01, 0x00, 0xfa, 0x01, 0x00},
	}

	for _, input := range inputs {
		output, err := Decode(input)
		if err != nil {
			t.Fatal(err)
		}
		roundTrip, err := Encode(output)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(input, roundTrip) {
			t.Errorf("Mismatch - Input: %# x RoundTripped: %# x", input, roundTrip)
		}
	}
}
