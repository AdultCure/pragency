package repository

type Authorization interface {

}

type Order interface {

}

type Repository struct {
	Authorization
	Order
}

func NewRepository() *Repository {
	return &Repository{}
}
