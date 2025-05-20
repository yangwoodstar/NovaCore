package modelStruct

type QuestionDetails struct {
	QuestionType   int      `json:"questionType"`
	Options        []string `json:"options"`
	Countdown      int      `json:"countdown"`
	Answer         []int    `json:"answers"`
	Describe       string   `json:"describe"`
	PublishAnswers bool     `json:"publishAnswers"`
}

type QuestionStart struct {
	ID              int             `json:"id"`
	QuestionDetails QuestionDetails `json:"questionDetails"`
}
