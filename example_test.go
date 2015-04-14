package xyz_test

import (
	"bytes"
	"fmt"
	"github.com/larzconwell/xyz"
	"io"
	"os"
)

func ExampleDecoder() {
	buf := bytes.NewBufferString(`5
Methane molecule
C  0.000000  0.000000  0.000000
H  0.000000  0.000000  1.089000
H  1.026719  0.000000 -0.363000
H -0.513360 -0.889165 -0.363000
H -0.513360  0.889165 -0.363000
3
Water molecule
O  0.00000 0.00000 0.00000
H  0.75700 0.58600 0.00000
H -0.75700 0.58600 0.00000`)

	decoder := xyz.NewDecoder(buf)
	for {
		molecule, err := decoder.Decode()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Println(molecule.Comment)
		for _, atom := range molecule.Atoms {
			fmt.Println(atom.Element, atom.X, atom.Y, atom.Z)
		}
	}
}

func ExampleEncoder() {
	molecule := &xyz.Molecule{
		Comment: "Water molecule",
		Atoms: []*xyz.Atom{
			{8, 0.00000, 0.00000, 0.00000},
			{1, 0.75700, 0.58600, 0.00000},
			{1, -0.075700, 0.58600, 0.00000},
		},
	}

	var buf bytes.Buffer
	encoder := xyz.NewEncoder(&buf)
	err := encoder.Encode(molecule)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, &buf)
}
