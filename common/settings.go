package common

type Settings struct {
	Theme    string `json:"theme"`
	Language string `json:"language"`
}

func NewDefaultSettings() *Settings {
	return &Settings{Theme: "light", Language: "en"}
}
