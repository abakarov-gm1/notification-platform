package handlers

import (
	"encoding/json"
	"net/http"
	"notification-service/internal/entity"
	"notification-service/internal/infrastructure/jwt"
)

type NotifUseCase interface {
	NotifCreate(newNotif *entity.NotificationEntity) error
	NotifDelete(notif *entity.NotificationEntity) error
}

type NotificationHandler struct {
	NotifUseCase NotifUseCase
}

func NewNotificationHandler(notifUseCase NotifUseCase) *NotificationHandler {
	return &NotificationHandler{NotifUseCase: notifUseCase}
}

func (r *NotificationHandler) NotifCreateHandler(w http.ResponseWriter, req *http.Request) {
	userId, er := jwt.DecodeToken(req.Header.Get("Authorization"))
	if er != nil {
		_, _ = w.Write([]byte(er.Error()))
	}

	var notif entity.NotificationEntity
	err := json.NewDecoder(req.Body).Decode(&notif)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	notif.UserID = userId

	if err := r.NotifUseCase.NotifCreate(&notif); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func (r *NotificationHandler) NotifDeleteHandler(w http.ResponseWriter, req *http.Request) {
	var notif entity.NotificationEntity
	err := json.NewDecoder(req.Body).Decode(&notif)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = r.NotifUseCase.NotifDelete(&notif)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
