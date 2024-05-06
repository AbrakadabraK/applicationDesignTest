package service

import "quest/internal/dto"

type OrderService interface {
	CreateOrder(order dto.Order, ordersStorage []dto.Order) error
}
