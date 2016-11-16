package lib

type Player struct {
	Name  string
	Map   *Map
	Ships [7]Ship
}

func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
		Map:  NewMap(),
		Ships: [7]Ship{
			Ship{Name: 'A', Size: 5},
			Ship{Name: 'B', Size: 4},
			Ship{Name: 'C', Size: 3},
			Ship{Name: 'D', Size: 2},
			Ship{Name: 'D', Size: 2},
			Ship{Name: 'S', Size: 1},
			Ship{Name: 'S', Size: 1},
		},
	}
}

func (p *Player) DropShip(x, y int, i int, d int) error {
	if i < 0 || i > len(p.Ships)-1 {
		return ErrWrongPos
	}
	p.Ships[i].Direction = d
	return p.Map.DropShip(x, y, p.Ships[i])
}
