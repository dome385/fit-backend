package models

type Übung struct {
	ID             string `json:"id"`
	Übung          string `json:"übung"`
	Sätze          int    `json:"sätze"`
	Wiederholungen int    `json:"wiederholungen"`
	Gewicht        int    `json:"gewicht"`
}
