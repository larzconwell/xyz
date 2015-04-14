package xyz

import (
	"bytes"
	"testing"
)

const caffeineMoleculeLit = `24
Caffeine molecule
H  -3.38041      -1.12724       5.73304E-01
N   9.66830E-01  -1.07374      -8.19823E-01
C   5.67293E-02   8.52719E-01   3.92316E-01
N  -1.37517      -1.02122      -5.70552E-02
C  -1.26150       2.59071E-01   5.23413E-01
C  -3.06834E-01  -1.68363      -7.16934E-01
C   1.13942       1.87412E-01  -2.70090E-01
N   5.60263E-01   2.08391       8.25159E-01
O  -4.92680E-01  -2.81806      -1.20947
C  -2.63281      -1.73040      -6.09530E-03
O  -2.23013       7.98862E-01   1.08997
H   2.54970       2.97350       6.22959E-01
C   2.05274      -1.73609      -1.49313
H  -2.48077      -2.72695       4.88263E-01
H  -3.00890      -1.90253      -1.04980
H   2.91761      -1.84815      -7.85787E-01
H   2.37879      -1.12119      -2.37437
H   1.71899      -2.74899      -1.84392
C  -1.51845E-01   3.09700       1.53483
C   1.89341       2.11812       4.19319E-01
N   2.28613       9.96844E-01  -2.44030E-01
H  -1.68703E-01   4.04366       9.30109E-01
H   3.53532E-01   3.29791       2.51777
H  -1.20745       2.75376       1.72030`

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
