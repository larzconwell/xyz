package xyz

import (
	"errors"
	"strconv"
	"strings"
)

var (
	// ErrElementSymbol occurs when the symbol is unrecognized.
	ErrElementSymbol = errors.New("xyz: element atomic symbol is unrecognized")
	// ErrElementNumber occurs when the given atomic number is out of range.
	ErrElementNumber = errors.New("xyz: element atomic number is out of range")
)

// Element is a single elements atomic number.
type Element int

// ParseElement parses the element from a slice of bytes. It recognizes both
// atomic numbers and atomic symbols.
func ParseElement(data []byte) (Element, error) {
	strdata := string(data)

	eli, err := strconv.Atoi(strdata)
	if err != nil {
		strdata = strings.ToLower(strdata)

		for i, name := range elementSymbols {
			if strings.ToLower(name) == strdata {
				return Element(i + 1), nil
			}
		}

		return 0, ErrElementSymbol
	}

	if eli > len(elementSymbols) || eli <= 0 {
		return 0, ErrElementNumber
	}

	return Element(eli), nil
}

// String returns the symbol for the element.
func (el Element) String() string {
	return elementSymbols[el-1]
}

var elementSymbols = [118]string{
	"H", "He", "Li", "Be", "B", "C", "N", "O", "F", "Ne",
	"Na", "Mg", "Al", "Si", "P", "S", "Cl", "Ar", "K", "Ca",
	"Sc", "Ti", "V", "Cr", "Mn", "Fe", "Co", "Ni", "Cu", "Zn",
	"Ga", "Ge", "As", "Se", "Br", "Kr", "Rb", "Sr", "Y", "Zr",
	"Nb", "Mo", "Tc", "Ru", "Rh", "Pd", "Ag", "Cd", "In", "Sn",
	"Sb", "Te", "I", "Xe", "Cs", "Ba", "La", "Ce", "Pr", "Nd",
	"Pm", "Sm", "Eu", "Gd", "Tb", "Dy", "Ho", "Er", "Tm", "Yb",
	"Lu", "Hf", "Ta", "W", "Re", "Os", "Ir", "Pt", "Au", "Hg",
	"Tl", "Pb", "Bi", "Po", "At", "Rn", "Fr", "Ra", "Ac", "Th",
	"Pa", "U", "Np", "Pu", "Am", "Cm", "Bk", "Cf", "Es", "Fm",
	"Md", "No", "Lr", "Rf", "Db", "Sg", "Bh", "Hs", "Mt", "Ds",
	"Rg", "Cn", "Uut", "Fl", "Uup", "Lv", "Uus", "Uuo",
}
