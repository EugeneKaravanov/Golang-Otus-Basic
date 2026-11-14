package chess

type Player struct {
	name string
}

func NewPlayer(name string) Player {
	return Player{
		name: name,
	}
}

func (player *Player) GetName() string {
	return player.name
}

func (player *Player) GetInfo() Player {
	return *player
}

func (player Player) Save() {}
