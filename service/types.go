package service

import (
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/go-chi/chi"
)

const (
	DefaultGWIncludeExpired = false
	DefaultGWStartAfter     = ""
	DefaultGWLimit          = ""
)

type Service interface {
	Name() string
	Router(*types.XnsContext, *chi.Mux) *chi.Mux
}
