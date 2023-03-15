package gamemodes

type LiteGamemode struct {
	Gamemode
}

func NewLiteGamemode() *LiteGamemode {
	return &LiteGamemode{}
}

func (c *LiteGamemode) GetName() string {
	return "Lite"
}

func (c *LiteGamemode) GetDescription() string {
	return "Lite gamemode"
}
