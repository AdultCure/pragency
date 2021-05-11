package backend

type Order struct {
	Id       int             `json:"id"`
	Category string          `json:"category"`
	Status   map[string]bool `json:"status"`
	Date     string          `json:"date"`
	Comment  string          `json:"comment"`
}

type OrderList struct {
	Id       int             `json:"id"`
	Category string          `json:"category"`
	Status   map[string]bool `json:"status"`
	Date     string          `json:"date"`
	Comment  string          `json:"comment"`
}
