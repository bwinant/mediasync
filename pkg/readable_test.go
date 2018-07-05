package pkg_test

import (
	ms "mediasync/pkg"
	"testing"
)

func TestByteLabel_Negative(t *testing.T) {
	e := "Error"
	l := ms.ByteLabel(-1);
	assertEquals(t, e, l)
}

func TestByteLabel_Zero(t *testing.T) {
	e := "0.00 B"
	l := ms.ByteLabel(0);
	assertEquals(t, e, l)
}

func TestByteLabel_B(t *testing.T) {
	e := "500.00 B"
	l := ms.ByteLabel(500);
	assertEquals(t, e, l)
}

func TestByteLabel_KB_Exact(t *testing.T) {
	e := "1.00 KB"
	l := ms.ByteLabel(1024);
	assertEquals(t, e, l)
}

func TestByteLabel_KB_Exact_Multiple(t *testing.T) {
	e := "3.00 KB"
	l := ms.ByteLabel(3072);
	assertEquals(t, e, l)
}

func TestByteLabel_KB_With_Fraction(t *testing.T) {
	e := "44.82 KB"
	l := ms.ByteLabel(45894);
	assertEquals(t, e, l)
}

func TestByteLabel_MB_Exact(t *testing.T) {
	e := "1.00 MB"
	l := ms.ByteLabel(1048576);
	assertEquals(t, e, l)
}

func TestByteLabel_MB_Exact_Multiple(t *testing.T) {
	e := "9.00 MB"
	l := ms.ByteLabel(9437184);
	assertEquals(t, e, l)
}

func TestByteLabel_MB_With_Fraction(t *testing.T) {
	e := "54.31 MB"
	l := ms.ByteLabel(56948265);
	assertEquals(t, e, l)
}

func TestByteLabel_GB_Exact(t *testing.T) {
	e := "1.00 GB"
	l := ms.ByteLabel(1073741824);
	assertEquals(t, e, l)
}

func TestByteLabel_GB_Exact_Multiple(t *testing.T) {
	e := "4.00 GB"
	l := ms.ByteLabel(4294967296);
	assertEquals(t, e, l)
}

func TestByteLabel_GB_With_Fraction(t *testing.T) {
	e := "7.91 GB"
	l := ms.ByteLabel(8489034596);
	assertEquals(t, e, l)
}

func TestByteLabel_TB_Exact(t *testing.T) {
	e := "1.00 TB"
	l := ms.ByteLabel(1099511627776);
	assertEquals(t, e, l)
}

func TestByteLabel_TB_Exact_Multiple(t *testing.T) {
	e := "2.00 TB"
	l := ms.ByteLabel(2.1990233e+12);
	assertEquals(t, e, l)
}

func TestByteLabel_TB_With_Fraction(t *testing.T) {
	e := "1.12 TB"
	l := ms.ByteLabel(1234567890000);
	assertEquals(t, e, l)
}

