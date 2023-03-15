package gamemodes

type ClassicGamemode struct {
	Gamemode
}

func NewClassicGamemode() *ClassicGamemode {
	return &ClassicGamemode{}
}

func (c *ClassicGamemode) GetName() string {
	return "Classic"
}

func (c *ClassicGamemode) GetDescription() string {
	return "Classic gamemode"
}
