package service

import (
	"fmt"
	"quest/internal/dto"
)

type Impl struct {
	RoomServ RoomAvailabilityService
}

func NewOrder(rs RoomAvailabilityService) *Impl {
	return &Impl{RoomServ: rs}
}

func (i *Impl) CreateOrder(order dto.Order, ordersStorage []dto.Order) error {
	// Мы проходили по всем дням в деапазоне бронирования и проверяли доступность каждого их них, неэффективно !
	// Сделал так : проверяем доступность бронирования всех номеров за период (это сократит сложность поиска)
	unavilableDays, err := i.RoomServ.CheckAvailability(order)
	if err != nil {
		return err
	}
	// используем список недуступных дней в логгере
	if len(unavilableDays) > 0 {
		return fmt.Errorf("hotel room is not available for selected dates: %v", unavilableDays)
	}
	// Бронирование будет доступно толлько тогда, когда доступны номера на все дни
	ordersStorage = append(ordersStorage, order)
	return nil
}
