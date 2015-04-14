package xyz

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

var (
	// ErrInvalidMolecule is returned when the given molecule input isn't valid.
	ErrInvalidMolecule = errors.New("xyz: input is not a valid molecule")
)

// Decoder reads and decodes lists of molecules.
type Decoder struct {
	reader  io.Reader
	scanner *bufio.Scanner
	err     error
}

// NewDecoder returns a decoder reading from the given reader.
//
// The decoder is buffered and may read data beyond the molecule data
// requests.
func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{reader: reader, scanner: bufio.NewScanner(reader)}
}

// Decode decodes the next molecule in the reader.
func (decoder *Decoder) Decode() (*Molecule, error) {
	if decoder.err != nil {
		return nil, decoder.err
	}
	molecule := &Molecule{Atoms: make([]*Atom, 0, 16)}

	num := 0
	i := -1
	for decoder.scanner.Scan() {
		line := strings.TrimSpace(decoder.scanner.Text())
		i++

		// The first line should contain the number of atoms to read.
		if i == 0 {
			var err error
			num, err = strconv.Atoi(line)
			if err != nil {
				decoder.err = ErrInvalidMolecule
				return nil, decoder.err
			}
			continue
		}

		// Next is the comment for the molecule.
		if i == 1 {
			molecule.Comment = line

			// If no molecules need to be read break.
			if num == 0 {
				break
			}

			continue
		}

		// Parse the atoms informations.
		items := strings.Fields(line)
		if len(items) != 4 {
			decoder.err = ErrInvalidMolecule
			return nil, decoder.err
		}

		element, err := ParseElement([]byte(items[0]))
		if err != nil {
			decoder.err = err
			return nil, err
		}

		var y float64
		var z float64
		x, err := strconv.ParseFloat(items[1], 64)
		if err == nil {
			y, err = strconv.ParseFloat(items[2], 64)
		}
		if err == nil {
			z, err = strconv.ParseFloat(items[3], 64)
		}
		if err != nil {
			decoder.err = err
			return nil, err
		}

		atom := &Atom{Element: element, X: x, Y: y, Z: z}
		molecule.Atoms = append(molecule.Atoms, atom)

		// If we're at the last atom for the molecule break, that way
		// we don't read items for another molecule.
		if i == num+1 {
			break
		}
	}

	// If i hasn't been set, then it's EOF.
	if i == -1 {
		decoder.err = io.EOF
		return nil, decoder.err
	}

	err := decoder.scanner.Err()
	if err != nil {
		decoder.err = err
		return nil, err
	}

	return molecule, nil
}
