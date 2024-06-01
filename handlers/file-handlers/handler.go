package fileHandlers

import (
	"ism/controllers/file-controllers"
)

type handler struct {
	service filecontrollers.Service
}

func NewCreateHandler(service filecontrollers.Service) *handler {
	return &handler{service: service}
}
