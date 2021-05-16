package backend

type Order struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
	Status   string `json:"status"`
	Date     string `json:"date"`
	Comment  string `json:"comment"`
	UserId   string `json:"user_id"`
}
