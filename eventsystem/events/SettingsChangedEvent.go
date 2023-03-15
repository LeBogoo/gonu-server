package events

type OnuSetting struct {
	Name     string   `json:"name"`
	Value    string   `json:"value"`
	Defaults []string `json:"defaults"`
}

type SettingsChangedEvent struct {
	BaseEvent
	Settings map[string]OnuSetting `json:"settings"`
}

func NewSettingsChangedEvent(settings map[string]OnuSetting) *SettingsChangedEvent {
	return &SettingsChangedEvent{
		BaseEvent: BaseEvent{Name: "SettingsChangedEvent"},
		Settings:  settings,
	}
}
