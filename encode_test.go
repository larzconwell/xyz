package xyz

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	var buf bytes.Buffer
	encoder := NewEncoder(&buf)
	err := encoder.Encode(caffeineMolecule)
	if err != nil {
		t.Fatal(err)
	}

	if buf.String() != caffeineMoleculeLit {
		t.Error("Encoded value does not match literal")
	}
}
