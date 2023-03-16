package onu

type UpdateAdminEvent struct {
	BaseEvent
	UUID string `json:"uuid"`
}

func NewUpdateAdminEvent(uuid string) *UpdateAdminEvent {
	return &UpdateAdminEvent{
		BaseEvent: BaseEvent{Name: "UpdateAdminEvent"},
		UUID:      uuid,
	}
}
