package handler

import "github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/service"

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}
