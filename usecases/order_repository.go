package usecases

import "github.com/thanchayawikgithub/learn-go-clean/entities"

type OrderRepository interface {
	Save(order entities.Order) error
}
