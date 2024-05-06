package service

import (
	"quest/internal/dto"
	"time"
)

type RoomServiceImpl struct{}

func NewRoom() *RoomServiceImpl {
	return &RoomServiceImpl{}
}

// Убрали map[time.Time]struct{} для отслеживания недоступных дней и используем и передаем в логер список недоступных дней
func (rs *RoomServiceImpl) CheckAvailability(order dto.Order) ([]time.Time, error) {
	roomAvailability := getAvailability()
	daysToBook := daysBetween(order.From, order.To)

	unavailableDays := make([]time.Time, 0)

	for _, dayToBook := range daysToBook {
		available := false
		for _, availability := range roomAvailability {
			if availability.Date.Equal(dayToBook) && availability.Quota > 0 {
				available = true
				break
			}
		}
		if !available {
			unavailableDays = append(unavailableDays, dayToBook)
		}
	}
	return unavailableDays, nil
}

func getAvailability() []dto.RoomAvailability {
	return []dto.RoomAvailability{
		{"reddison", "lux", date(2024, 1, 1), 1},
		{"reddison", "lux", date(2024, 1, 2), 1},
		{"reddison", "lux", date(2024, 1, 3), 0},
		{"reddison", "lux", date(2024, 1, 4), 1},
		{"reddison", "lux", date(2024, 1, 5), 0},
	}
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func daysBetween(from time.Time, to time.Time) []time.Time {
	if from.After(to) {
		return nil
	}

	days := make([]time.Time, 0)
	for d := toDay(from); !d.After(toDay(to)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days
}

func toDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}
