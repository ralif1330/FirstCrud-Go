package server

import (
	"database/sql"
	"log"
	"portfolioghOne/controller"
	"portfolioghOne/repositories"
	"portfolioghOne/service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	ControllerManager controller.ControllersManager
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {

	repositoriesManager := repositories.NewRepositoriesManager(dbHandler)

	servicesManager := service.NewServiceManager(repositoriesManager)

	controllerManager := controller.NewControllersManager(servicesManager)

	router := gin.Default()

	InitRouter(router, controllerManager)

	return HttpServer{
		config:            config,
		router:            router,
		ControllerManager: *controllerManager,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))

	if err != nil {
		log.Fatalf("Error while starting HTTP Server: %v", err)
	}
}
