package models
type Question struct {
	Question  string `json:"question"`
	//MaxLength int    `json:"max_length"`
}

type AnswerResponse struct {
	Answer string `json:"answer"`
}