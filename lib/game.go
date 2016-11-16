package lib

type Game struct {
	players [2]*Player
	holder  int
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Join(p *Player) {
	if g.players[0] == nil {
		g.players[0] = p
		return
	}
	g.players[1] = p
}

func (g *Game) Shoot(x, y int) {
	g.players[g.holder].Map.Shoot(x, y)
	g.holder = g.holder // TODO change player
}
