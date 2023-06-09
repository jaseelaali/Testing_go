//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/jaseelali/go-gin-clean-arch/pkg/api"
	handler "github.com/jaseelali/go-gin-clean-arch/pkg/api/handler"
	config "github.com/jaseelali/go-gin-clean-arch/pkg/config"
	db "github.com/jaseelali/go-gin-clean-arch/pkg/db"
	repository "github.com/jaseelali/go-gin-clean-arch/pkg/repository"
	usecase "github.com/jaseelali/go-gin-clean-arch/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler, http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
