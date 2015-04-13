package xyz

import (
	"testing"
)

func TestParseElementInt(t *testing.T) {
	e, err := ParseElement([]byte("5"))
	if err != nil {
		t.Fatal(err)
	}

	if e != 5 {
		t.Error("Atomic number returned isn't expected")
	}

	if e.String() != "B" {
		t.Error("Atomic symbol returned isn't expected")
	}
}

func TestParseElementInt2(t *testing.T) {
	e, err := ParseElement([]byte("76"))
	if err != nil {
		t.Fatal(err)
	}

	if e != 76 {
		t.Error("Atomic number returned isn't expected")
	}

	if e.String() != "Os" {
		t.Error("Atomic symbol returned isn't expected")
	}
}

func TestParseElementOOB(t *testing.T) {
	_, err := ParseElement([]byte("10000"))
	if err == nil {
		t.Error("Positive out of bounds parse succeeded when it should've failed")
	}

	_, err = ParseElement([]byte("-500"))
	if err == nil {
		t.Error("Negative out of bounds parse succeeded when it should've failed")
	}
}

func TestParseElementStr(t *testing.T) {
	e, err := ParseElement([]byte("ru"))
	if err != nil {
		t.Fatal(err)
	}

	if e != 44 {
		t.Error("Atomic number returned isn't expected")
	}

	if e.String() != "Ru" {
		t.Error("Atomic symbol returned isn't expected")
	}
}

func TestParseElementStr2(t *testing.T) {
	e, err := ParseElement([]byte("FE"))
	if err != nil {
		t.Fatal(err)
	}

	if e != 26 {
		t.Error("Atomic number returned isn't expected")
	}

	if e.String() != "Fe" {
		t.Error("Atomic symbol returned isn't expected")
	}
}

func TestParseElementUnrecognized(t *testing.T) {
	_, err := ParseElement([]byte("larz"))
	if err == nil {
		t.Error("Unrecognized atomic symbol succeeded when it should've failed")
	}
}
