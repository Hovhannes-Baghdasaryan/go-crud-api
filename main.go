package main

import (
	"crud-go-api/config"
	router "crud-go-api/internal/router/tags"
	constant "crud-go-api/libs/common/constant/logger"
	helper "crud-go-api/libs/common/helper/error"
	logger "crud-go-api/libs/common/logger/main"
	common "crud-go-api/libs/common/router"
	configuration "crud-go-api/libs/data-layer/configuration/ent"
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	cfg := config.ConfigLoad()

	logger.LogInfo(logger.LoggerPayload{Message: slog.String("env", cfg.Env).String(), FuncName: constant.MainBoostrap})

	clientDB := configuration.DatabaseConnection()

	baseRoute := common.BaseRouter()
	routes := router.InjectTagRouter(baseRoute, clientDB)

	logger.LogInfo(logger.LoggerPayload{FuncName: constant.MainBoostrap, Message: fmt.Sprintf("Starting Server %s", slog.String("address", cfg.Address))})

	server := &http.Server{
		Addr:    cfg.Address,
		Handler: *routes,
	}

	if err := server.ListenAndServe(); err != nil {
		helper.PanicIfError(err)
	}
}
