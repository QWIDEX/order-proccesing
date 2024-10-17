package main

import "context"

type db struct {
	//
}

type OrdersDb interface {
	NewOrder() error
}

func NewOrdersDB(ctx context.Context) OrdersDb {
	return &db{}
}

func (db *db) NewOrder() error {
	return nil
}
