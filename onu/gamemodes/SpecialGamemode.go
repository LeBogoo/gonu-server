package gamemodes

type SpecialGamemode struct {
	Gamemode
}

func NewSpecialGamemode() *SpecialGamemode {
	return &SpecialGamemode{}
}

func (c *SpecialGamemode) GetName() string {
	return "Special"
}

func (c *SpecialGamemode) GetDescription() string {
	return "Special gamemode"
}
