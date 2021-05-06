package checkersLogic

// player interface is prototype for player

type _player struct {
	pieces    [12]*_piece
	remaining int
	isAI      bool
}

// throws error if asset is missing

func (p _player) Pieces() []*_piece {
	return p.pieces[:]
}

// initilizes player class
func player_init(p, dir int, ai bool) _player {
	temp := _player{
		pieces:    [12]*_piece{},
		remaining: 12,
		isAI:      ai,
	}

	for i := 0; i < len(temp.pieces); i++ {
		temp.pieces[i] = &_piece{Owner: p, direction: dir}
	}

	return temp
}
