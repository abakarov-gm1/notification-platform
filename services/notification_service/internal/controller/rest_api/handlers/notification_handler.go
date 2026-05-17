package handlers

import (
	"encoding/json"
	"net/http"
	"notification-service/internal/entity"
	"notification-service/internal/infrastructure/jwt"
)

type NotifyUseCase interface {
	NotifyCreate(newNotify *entity.NotificationEntity) error
	NotifyDelete(notify *entity.NotificationEntity) error
}

type NotificationHandler struct {
	NotifyUseCase NotifyUseCase
}

func NewNotificationHandler(notifyUseCase NotifyUseCase) *NotificationHandler {
	return &NotificationHandler{NotifyUseCase: notifyUseCase}
}

func (r *NotificationHandler) NotifyCreateHandler(w http.ResponseWriter, req *http.Request) {
	userId, er := jwt.DecodeToken(req.Header.Get("Authorization"))
	if er != nil {
		_, _ = w.Write([]byte(er.Error()))
	}

	var notify entity.NotificationEntity
	err := json.NewDecoder(req.Body).Decode(&notify)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	notify.UserID = userId

	if err := r.NotifyUseCase.NotifyCreate(&notify); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func (r *NotificationHandler) NotifyDeleteHandler(w http.ResponseWriter, req *http.Request) {
	var notify entity.NotificationEntity
	err := json.NewDecoder(req.Body).Decode(&notify)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = r.NotifyUseCase.NotifyDelete(&notify)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
