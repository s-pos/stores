package main

import (
	"spos/stores/controllers"
	"spos/stores/repository"
	"spos/stores/routes"
	"spos/stores/usecase"

	"github.com/s-pos/go-utils/adapter"
	"github.com/s-pos/go-utils/config"
	"github.com/s-pos/go-utils/middleware"
	"github.com/s-pos/go-utils/utils/server"
)

func init() {
	serviceName := "stores"

	config.Load(serviceName)
}

func main() {
	log := config.Logrus()
	timezone := config.Timezone()

	db := adapter.DBConnection()
	redis := adapter.GetClientRedis()

	middleware := middleware.NewMiddleware(redis, log, timezone)

	repo := repository.NewRepository(db)
	uc := usecase.NewUsecase(repo)
	ctrl := controllers.NewController(uc)

	router := routes.NewRoute(middleware, ctrl)

	log.Fatal(server.Wrapper(router.Router()))
}
