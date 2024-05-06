package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"quest/internal/dto"
	"quest/internal/pkg/logger"
	"quest/internal/service"
)

var Orders []dto.Order

type HTTPHandler struct {
	OrderService            service.OrderService
	RoomAvailabilityService service.RoomAvailabilityService
}

func NewHTTPHandler(os service.OrderService, ra service.RoomAvailabilityService) *HTTPHandler {
	return &HTTPHandler{
		OrderService:            os,
		RoomAvailabilityService: ra,
	}
}

// Была неатомарнасть операций , т.е если бронирование не доступно все равно создавалось бронирование.
func (h *HTTPHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder dto.Order
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		logger.LogErrorf("error decode request order:  %v", newOrder)
		return
	}
	// прокидываю ошибку и текст при невозможности бронирование
	err = h.OrderService.CreateOrder(newOrder, Orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		logger.LogErrorf("error create order:  %v", newOrder)
		return
	}
	fmt.Println(Orders)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newOrder)
	if err != nil {
		logger.LogErrorf("error Encode order:  %v", newOrder)
		return
	}
}
