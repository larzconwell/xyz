package xyz

import (
	"io"
	"strings"
	"testing"
)

var caffeineMolecule = &Molecule{
	Comment: "Caffeine molecule",
	Atoms: []*Atom{
		{1, -3.38041, -1.12724, 5.73304E-01},
		{7, 9.66830E-01, -1.07374, -8.19823E-01},
		{6, 5.67293E-02, 8.52719E-01, 3.92316E-01},
		{7, -1.37517, -1.02122, -5.70552E-02},
		{6, -1.26150, 2.59071E-01, 5.23413E-01},
		{6, -3.06834E-01, -1.68363, -7.16934E-01},
		{6, 1.13942, 1.87412E-01, -2.70090E-01},
		{7, 5.60263E-01, 2.08391, 8.25159E-01},
		{8, -4.92680E-01, -2.81806, -1.20947},
		{6, -2.63281, -1.73040, -6.09530E-03},
		{8, -2.23013, 7.98862E-01, 1.08997},
		{1, 2.54970, 2.97350, 6.22959E-01},
		{6, 2.05274, -1.73609, -1.49313},
		{1, -2.48077, -2.72695, 4.88263E-01},
		{1, -3.00890, -1.90253, -1.04980},
		{1, 2.91761, -1.84815, -7.85787E-01},
		{1, 2.37879, -1.12119, -2.37437},
		{1, 1.71899, -2.74899, -1.84392},
		{6, -1.51845E-01, 3.09700, 1.53483},
		{6, 1.89341, 2.11812, 4.19319E-01},
		{7, 2.28613, 9.96844E-01, -2.44030E-01},
		{1, -1.68703E-01, 4.04366, 9.30109E-01},
		{1, 3.53532E-01, 3.29791, 2.51777},
		{1, -1.20745, 2.75376, 1.72030},
	},
}

const caffeineMoleculeMessyLit = `24
     Caffeine molecule
     1  -3.38041 -1.12724       5.73304E-01
N   9.66830E-01  -1.07374      -8.19823E-01
C   5.67293E-02   8.52719E-01   3.92316E-01
N  -1.37517      -1.02122             -5.70552E-02
C  -1.26150       2.59071E-01   5.23413E-01
7  -3.06834E-01       -1.68363      -7.16934E-01
C   1.13942       1.87412E-01  -2.70090E-01
      N   5.60263E-01   2.08391       8.25159E-01
8  -4.92680E-01  -2.81806      -1.20947
C  -2.63281      -1.73040      -6.09530E-03
    O  -2.23013       7.98862E-01   1.08997
H   2.54970       2.97350       6.22959E-01
C   2.05274      -1.73609      -1.49313
H  -2.48077      -2.72695       4.88263E-01
1  -3.00890      -1.90253      -1.04980
1        2.91761      -1.84815      -7.85787E-01
H   2.37879      -1.12119      -2.37437
    H   1.71899      -2.74899      -1.84392
6  -1.51845E-01   3.09700       1.53483
C   1.89341       2.11812       4.19319E-01
N   2.28613           9.96844E-01  -2.44030E-01
H  -1.68703E-01        4.04366       9.30109E-01
H 3.53532E-01 3.29791 2.51777
H  -1.20745       2.75376       1.72030`

const multiMoleculeLit = `0
Nothing

5
Methane
C  0.000000  0.000000  0.000000
H  0.000000  0.000000  1.089000
H  1.026719  0.000000 -0.363000
H -0.513360 -0.889165 -0.363000
H -0.513360  0.889165 -0.363000

3
Water
O  0.00000 0.00000 0.00000
H  0.75700 0.58600 0.00000
H -0.75700 0.58600 0.00000

`

const invalidARContent = `!<arch>
some-file-name    1257894000 ...`

func TestDecoderLiteral(t *testing.T) {
	decoder := NewDecoder(strings.NewReader(caffeineMoleculeLit))
	molecule, err := decoder.Decode()
	if err != nil {
		t.Fatal(err)
	}

	if molecule.Comment != caffeineMolecule.Comment {
		t.Error("Comment doesn't match source")
	}

	if len(molecule.Atoms) != len(caffeineMolecule.Atoms) {
		t.Error("Number of atoms doesn't match source")
	}

	_, err = decoder.Decode()
	if err == nil || err != io.EOF {
		t.Error("Decode exhausted reader returns incorrect error")
	}
}

func TestDecoderGenerated(t *testing.T) {
	decoder := NewDecoder(strings.NewReader(caffeineMolecule.String()))
	molecule, err := decoder.Decode()
	if err != nil {
		t.Fatal(err)
	}

	if molecule.Comment != caffeineMolecule.Comment {
		t.Error("Comment doesn't match source")
	}

	if len(molecule.Atoms) != len(caffeineMolecule.Atoms) {
		t.Error("Number of atoms doesn't match source")
	}
}

func TestDecoderMutlti(t *testing.T) {
	decoder := NewDecoder(strings.NewReader(multiMoleculeLit))
	for {
		molecule, err := decoder.Decode()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}

		switch molecule.Comment {
		case "Nothing":
			if len(molecule.Atoms) != 0 {
				t.Error("Nothing molecule has incorrect number of atoms")
			}
		case "Methane":
			if len(molecule.Atoms) != 5 {
				t.Error("Methane molecule has incorrect number of atoms")
			}
		case "Water":
			if len(molecule.Atoms) != 3 {
				t.Error("Water molecule has incorrect number of atoms")
			}
		default:
			t.Error("Unknown comment for molecule")
		}
	}
}

func TestDecoderInvalid(t *testing.T) {
	decoder := NewDecoder(strings.NewReader(invalidARContent))
	_, err := decoder.Decode()
	if err == nil {
		t.Error("Decoder succeeded with invalid input.")
	}
}
