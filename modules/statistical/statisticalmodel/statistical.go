package statisticalmodel

type StatisticalElement struct {
	TaskCount      int `json:"task_count"`
	HabitCount     int `json:"habit_count"`
	UserCount      int `json:"user_count"`
	ChallengeCount int `json:"challenge_count"`
}

type Statistical struct {
	TaskCount      int                 `json:"task_count"`
	HabitCount     int                 `json:"habit_count"`
	UserCount      int                 `json:"user_count"`
	ChallengeCount int                 `json:"challenge_count"`
	Time           string              `json:"time"`
	Element        *StatisticalElement `json:"element"`
}

type Filter struct {
	Time string `json:"time" form:"time"`
}

func NewStatisticalElement(tc, hc, uc, cc int) *StatisticalElement {
	return &StatisticalElement{
		TaskCount:      tc,
		HabitCount:     hc,
		UserCount:      uc,
		ChallengeCount: cc,
	}
}

func NewStatistical(tc, hc, uc, cc int, time string, element *StatisticalElement) *Statistical {
	return &Statistical{
		TaskCount:      tc,
		HabitCount:     hc,
		UserCount:      uc,
		ChallengeCount: cc,
		Time:           time,
		Element:        element,
	}
}
