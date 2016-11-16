package main

import (
	"fmt"
	"math/rand"
	"os"
)

type Ship struct {
	name  string
	size  int
	score int
}

var ships = [...]Ship{
	{"Aircraft Carrier", 5, 20},
	{"Battleship", 4, 12},
	{"Submarine", 3, 6},
	{"Destroyer", 3, 6},
	{"Cruiser", 3, 6},
	{"Patrol Boat", 2, 2},
}

const (
	boardSize      = 16
	guessesPerTurn = 5
	maxTurns       = 6
)

func main() {
	board := genBoard()

	won := false
game:
	for turn := 0; turn < maxTurns; turn++ {
		fmt.Printf("Turn #%v\n", turn+1)
		var guesses [guessesPerTurn]int
		for i := range guesses {
			guess := getGuess(i + 1)
			if guess < 0 {
				fmt.Println("invalid coordinate")
				os.Exit(1)
			}
			guesses[i] = guess
		}
		hits := 0
		for _, guess := range guesses {
		check:
			for j, ship := range board.shipSpaces {
				for i, coord := range ship {
					if guess == coord {
						x, y := coord>>8, coord&0x0f
						fmt.Printf("Hit at (%v, %v)!\n", x, y)
						ship[i] = ship[len(ship)-1]
						board.shipSpaces[j] = ship[:len(ship)-1]
						if len(board.shipSpaces[j]) == 0 {
							fmt.Printf("You sunk the %s!\n", board.ships[j].name)
						}
						board.spaces--
						hits++
						if board.spaces == 0 {
							won = true
							break game
						}
						break check
					}
				}
			}
		}
		if hits == 0 {
			fmt.Println("No hits!")
		}
	}
	if !won {
		fmt.Println("You lost!")
	}
}

func getGuess(guessNum int) int {
	fmt.Printf("Enter x for guess #%d: ", guessNum)
	var x int
	fmt.Scanln(&x)
	fmt.Printf("Enter y for guess #%d: ", guessNum)
	var y int
	fmt.Scanln(&y)
	return x<<8 | y
}

type board struct {
	spaces     int
	ships      [5]Ship
	shipSpaces [5][]int
}

func genBoard() board {
	rng := rand.New(rand.NewSource(0))

	// choose which ships we're going to place
	ships := append([]Ship(nil), ships[:]...)
	// shuffle ships
	for i := range ships {
		j := rng.Intn(i + 1)
		ships[i], ships[j] = ships[j], ships[i]
	}
	// cut off last one
	ships = ships[:5]

	var b board
	for k, ship := range ships {
	place:
		for {
			// rotate == true means that the ship is vertical
			rotate := rand.Int()&1 != 0
			maxx := boardSize
			maxy := boardSize
			if rotate {
				maxy -= ship.size
			} else {
				maxx -= ship.size
			}
			x := rand.Intn(maxx)
			y := rand.Intn(maxy)
			coord := x<<8 | y
			coords := make([]int, ship.size)
			for i := range coords {
				coords[i] = coord
				if rotate {
					coord++
				} else {
					coord += 1 << 8
				}
			}
			// see if any of the coords are the same as the coords that
			// are already occupied
			for _, coord := range coords {
				for _, placedShip := range b.shipSpaces {
					for _, coord2 := range placedShip {
						if coord == coord2 {
							continue place
						}
					}
				}
			}
			//fmt.Printf("Placed %s at (%v, %v); rot: %v\n", ship.name, x, y, rotate)
			b.spaces += len(coords)
			b.ships[k] = ship
			b.shipSpaces[k] = coords
			break
		}
	}
	return b
}
