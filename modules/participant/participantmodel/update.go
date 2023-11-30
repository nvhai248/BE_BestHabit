package participantmodel

type ParticipantUpdate struct {
	Status *string `json:"status" db:"status"`
}

func (ParticipantUpdate) TableName() string {
	return Participant{}.TableName()
}
