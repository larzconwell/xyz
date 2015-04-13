package xyz

import (
	"io"
)

// Encoder writes and encodes lists of molecules.
type Encoder struct {
	writer    io.Writer
	leadingnl bool
}

// NewEncoder creates an encoder writing to writer.
func NewEncoder(writer io.Writer) *Encoder {
	return &Encoder{writer: writer}
}

// Encode encodes the given molecule writing to the writer.
func (encoder *Encoder) Encode(molecule *Molecule) error {
	str := molecule.String()
	if encoder.leadingnl {
		str = "\n" + str
	}

	_, err := encoder.writer.Write([]byte(str))
	if err != nil {
		return err
	}

	encoder.leadingnl = true
	return nil
}
