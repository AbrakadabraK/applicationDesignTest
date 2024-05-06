package service

import (
	"quest/internal/dto"
	"time"
)

type RoomAvailabilityService interface {
	CheckAvailability(order dto.Order) ([]time.Time, error)
}
