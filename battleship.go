package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Pos struct {
	IsOccupied bool
	IsShot     bool
}

func (p *Pos) Out() string {
	if p.IsShot && p.IsOccupied {
		return "#"
	}
	if p.IsOccupied {
		return "O"
	}
	if p.IsShot {

		return "_"
	}
	return " "
}

var len = 64

func main() {
	board := []*Pos{}
	for i := 0; i < len; i++ {
		is := i%6 == 0
		board = append(board, &Pos{IsOccupied: is})
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		param1 := r.URL.Query().Get("b")
		if param1 != "" {
			coordinate, err := strconv.Atoi(param1)
			if err != nil {
				fmt.Fprintf(w, "%s", err)
				return
			}
			board[coordinate].IsShot = true
		}
		log.Printf("b = %s", param1)
		for i, p := range board {
			fmt.Fprintf(w, "%s", p.Out())
			if i%8 == 7 {
				fmt.Fprintf(w, "\n")
			}
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
