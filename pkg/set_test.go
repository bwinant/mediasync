package pkg_test

import (
	ms "mediasync/pkg"
	"testing"
)

func TestSet_Add(t *testing.T) {
	s := ms.NewSet()
	s.Add("a")

	assertFalse(t, s.Contains("y"))
	assertTrue(t, s.Contains("a"))
}

func TestSet_AddAll(t *testing.T) {
	s := ms.NewSet()

	a := [3]string{"x", "y", "z"}
	s.AddAll(a[:])

	assertFalse(t, s.Contains("a"))
	assertTrue(t, s.Contains("y"))
}

func TestSet_Remove(t *testing.T) {
	s := ms.NewSet()

	s.Add("a")
	assertTrue(t, s.Contains("a"))

	s.Remove("a")
	assertFalse(t, s.Contains("a"))
}

func TestSet_Size(t *testing.T) {
	s := ms.NewSet()
	assertTrue(t, s.Size() == 0)

	s.Add("a")
	assertTrue(t, s.Size() == 1)

	s.Remove("a")
	assertTrue(t, s.Size() == 0)

	a := [3]string{"x", "y", "z"}
	s.AddAll(a[:])
	assertTrue(t, s.Size() == 3)

	s.Add("a")
	assertTrue(t, s.Size() == 4)

	s.Add("a")
	assertTrue(t, s.Size() == 4)
}
