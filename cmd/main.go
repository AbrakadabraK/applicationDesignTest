// Ниже реализован сервис бронирования номеров в отеле. В предметной области
// выделены два понятия: Order — заказ, который включает в себя даты бронирования
// и контакты пользователя, и RoomAvailability — количество свободных номеров на
// конкретный день.
//
// Задание:
// - провести рефакторинг кода с выделением слоев и абстракций
// - применить best-practices там где это имеет смысл
// - исправить имеющиеся в реализации логические и технические ошибки и неточности
package main

import (
	"os"
	"quest/internal/pkg/app"
	"quest/internal/pkg/logger"
)

/*
Разбил сервис на слои :

Прописал контроллер( по хорошему наверное еще Миделвейр использовать) - тут лежит HTTPHandler который обрабатывает запросы по эндпоинту /order

# Прописал package app и засунул в него app и logger

app.go - запуск всех наших сервисов хендлеров и подрнятие сервиса по адресу
logger.go - просто пакет своего логгера

internal - наша бизнес логика
dto - вынес сюда энтити проекта Order и AvailRoom
service  - бизнес логика (менеджер)
order.go и order_impl.go - положил интерфейсы OrderService с методом CreateOrder и в impl реализовал создание заказа

room.go и room_impl.go - интерфейс RoomAvailabilityService и его метод CheckAvailability
в impl - реализовал его
по хорошему думаю интерфейсы использовать в контроллере - потому что там мы их используем ...
*/
func main() {
	a, err := app.New()
	if err != nil {
		logger.LogErrorf("Error create application")
		os.Exit(1)
	}
	err = a.Run()
	if err != nil {
		logger.LogErrorf("Server failed: %s", err)
		os.Exit(1)
	}
}

//func createOrder(w http.ResponseWriter, r *http.Request) {
//	var newOrder dto.Order
//	json.NewDecoder(r.Body).Decode(&newOrder)
//
//	daysToBook := daysBetween(newOrder.From, newOrder.To)
//
//	unavailableDays := make(map[time.Time]struct{})
//	for _, day := range daysToBook {
//		unavailableDays[day] = struct{}{}
//	}
//
//	for _, dayToBook := range daysToBook {
//		for i, availability := range Availability {
//			if !availability.Date.Equal(dayToBook) || availability.Quota < 1 {
//				continue
//			}
//			availability.Quota -= 1
//			Availability[i] = availability
//			delete(unavailableDays, dayToBook)
//		}
//	}
//
//	if len(unavailableDays) != 0 {
//		http.Error(w, "Hotel room is not available for selected dates", http.StatusInternalServerError)
//		LogErrorf("Hotel room is not available for selected dates:\n%v\n%v", newOrder, unavailableDays)
//		return
//	}
//
//	Orders = append(Orders, newOrder)
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusCreated)
//	err := json.NewEncoder(w).Encode(newOrder)
//	if err != nil {
//		logger.LogErrorf("Error Encode New Order:\n%v", newOrder)
//		return
//	}
//
//	logger.LogInfo("Order successfully created: %v", newOrder)
//}
