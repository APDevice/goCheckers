package checkersLogic

// player interface is prototype for player

type player struct {
	pieces    [12]*Piece
	remaining int
	isAI      bool
}

// throws error if asset is missing

func (p player) Pieces() []*Piece {
	return p.pieces[:]
}

// initilizes player class
func player_init(p, dir int, ai bool) player {
	temp := player{
		pieces:    [12]*Piece{},
		remaining: 12,
		isAI:      ai,
	}

	for i := 0; i < len(temp.pieces); i++ {
		temp.pieces[i] = &Piece{Owner: p, direction: dir}
	}

	return temp
}
