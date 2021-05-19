package backend

type Order struct {
	Id       int    `json:"id"`
	Category string `json:"category" binding:"required"`
	Status   string `json:"status"`
	Date     string `json:"date" binding:"required"`
	Comment  string `json:"comment"`
	UserId   string `json:"user_id"`
}
