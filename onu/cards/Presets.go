package cards

func NewClassicPreset() *Preset {
	return &Preset{
		Colors: []Color{*ColorFrom("r"), *ColorFrom("g"), *ColorFrom("b"), *ColorFrom("y")},
		Types:  []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
	}
}

func NewWishPreset() *Preset {
	return &Preset{
		Colors: []Color{*ColorFrom("none")},
		Types:  []string{"w", "p4"},
	}
}
