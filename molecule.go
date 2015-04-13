package xyz

import (
	"bytes"
	"fmt"
	"strconv"
	"text/tabwriter"
)

// Molecule contains the atoms that make it up.
type Molecule struct {
	Comment string
	Atoms   []*Atom
}

// String prints the molecule in its literal format.
func (molecule *Molecule) String() string {
	str := strconv.Itoa(len(molecule.Atoms)) + "\n"
	str += molecule.Comment + "\n"

	var buf bytes.Buffer
	writer := tabwriter.NewWriter(&buf, 2, 2, 2, ' ', 0)
	for _, atom := range molecule.Atoms {
		writer.Write([]byte(atom.String() + "\n"))
	}
	writer.Flush()
	str += buf.String()

	return str[:len(str)-1]
}

// Atom contains the element and it's cartesian coords.
type Atom struct {
	Element Element
	X       float64
	Y       float64
	Z       float64
}

// String prints the atom in its literal form.
func (atom *Atom) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s", atom.Element.String(),
		formatCoord(atom.X), formatCoord(atom.Y), formatCoord(atom.Z))
}

func formatCoord(coord float64) string {
	str := strconv.FormatFloat(coord, 'E', 5, 64)

	// If no exponent then just remove it.
	idx := len(str) - 4
	if str[idx:] == "E+00" {
		str = str[:idx]
	}

	// Align better with negative numbers.
	if coord >= 0 {
		str = " " + str
	}

	return str
}
