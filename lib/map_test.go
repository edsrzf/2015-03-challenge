package lib

import (
	"testing"
)

func printMap(t *testing.T, m *Map) {
	for _, row := range m {
		t.Log(row)
	}
}

func TestDropShip(t *testing.T) {
	m := NewMap()
	if err := m.DropShip(-1, 0, Ship{Name: 'A', Size: 5, Direction: V}); err != ErrWrongPos {
		t.Fatalf("ErrWrongPos: %s", err)
	}
	if err := m.DropShip(9, 0, Ship{Name: 'A', Size: 5, Direction: V}); err != ErrWrongPos {
		t.Fatalf("ErrWrongPos: %s", err)
	}
	if err := m.DropShip(0, 0, Ship{Name: 'A', Size: 5, Direction: V}); err != nil {
		t.Fatal(err)
	}
	if err := m.DropShip(1, 1, Ship{Name: 'A', Size: 5, Direction: V}); err != nil {
		t.Fatal(err)
	}
	if err := m.DropShip(0, 0, Ship{Name: 'A', Size: 5, Direction: V}); err != ErrWrongPos {
		t.Fatalf("ErrWrongPos: %s", err)
	}
	printMap(t, m)
}

func TestShootShip(t *testing.T) {
	m := NewMap()
	if err := m.DropShip(0, 0, Ship{Name: 'A', Size: 5, Direction: V}); err != nil {
		t.Fatal(err)
	}
	if err := m.Shoot(0, 0); err != nil {
		t.Fatal(err)
	}
	if err := m.Shoot(1, 1); err != ErrMiss {
		t.Fatalf("ErrMiss needed: %s", err)
	}
	if err := m.Shoot(0, 0); err != ErrWrongPos {
		t.Fatalf("ErrWrongPos needed: %s", err)
	}
	printMap(t, m)
}
