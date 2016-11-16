package lib

import (
	"fmt"
)

const (
	H = iota
	V
)

var (
	ErrWrongPos = fmt.Errorf("Wrong Position")
	ErrMiss     = fmt.Errorf("Missed")
)

type Map [8][8]rune // X,Y

func NewMap() *Map {
	var m Map
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			m[x][y] = 'O'
		}
	}
	return &m
}

func (m *Map) DropShip(x, y int, s Ship) error {
	if x < 0 || y < 0 || x > 8 || y > 8 {
		return ErrWrongPos
	}
	if s.Direction == H {
		if (x + s.Size) >= 8 {
			return ErrWrongPos
		}
		for i := 0; i < s.Size; i++ {
			if m[x+i][y] != 'O' {
				return ErrWrongPos
			}
			m[x+i][y] = s.Name
		}
		return nil
	}
	if (y + s.Size) >= 8 {
		return ErrWrongPos
	}
	for i := 0; i < s.Size; i++ {
		if m[x][y+i] != 'O' {
			return ErrWrongPos
		}
		m[x][y+i] = s.Name
	}
	return nil
}

func (m *Map) Shoot(x, y int) error {
	switch m[x][y] {
	case 'O':
		return ErrMiss
	case 'X':
		return ErrWrongPos
	}
	m[x][y] = 'X'
	return nil
}

type Ship struct {
	Name      rune
	Size      int
	Direction int
}
