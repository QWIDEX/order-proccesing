package main

type service struct {
	db OrdersDb
}

type Service interface {
	CreateOrder() error
}

func NewService(db OrdersDb) Service {
	return &service{db: db}
}

func (s *service) CreateOrder() error {
	return nil
}
