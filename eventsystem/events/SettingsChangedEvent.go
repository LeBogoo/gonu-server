package events

type OnuSettings struct {
	Name     string   `json:"name"`
	Value    string   `json:"value"`
	Defaults []string `json:"defaults"`
}

type SettingsChangedEvent struct {
	BaseEvent
	Settings map[string]OnuSettings `json:"settings"`
}

func NewSettingsChangedEvent(settings map[string]OnuSettings) *SettingsChangedEvent {
	return &SettingsChangedEvent{
		BaseEvent: BaseEvent{Name: "SettingsChangedEvent"},
		Settings:  settings,
	}
}
