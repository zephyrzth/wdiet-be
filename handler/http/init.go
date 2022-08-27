package http

import "github.com/zephyrzth/wdiet-be/usecase"

type handler struct {
	usecase usecase.UsecaseInterface
}

func New(usecase usecase.UsecaseInterface) *handler {
	return &handler{
		usecase: usecase,
	}
}
