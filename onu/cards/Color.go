package cards

type Color struct {
	Color string `json:"color"`
}

func (c *Color) String() string {

	return c.Color
}

func (c *Color) Equals(other *Color) bool {
	return c.Color == other.Color
}

func (c *Color) IsAny(colors ...string) bool {
	for _, color := range colors {
		if c.Color == color {
			return true
		}
	}
	return false
}

func ColorFrom(s string) *Color {
	return &Color{
		Color: s,
	}
}
