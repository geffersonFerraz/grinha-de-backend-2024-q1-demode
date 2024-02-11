package main

import (
	"context"
	"fmt"

	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-demode/config"
	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-demode/database"
	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-demode/src/controller"
	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-demode/src/repository"
	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-demode/src/server"
	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-demode/src/usecase"
)

func main() {
	// database -> repository -> usecase -> controller -> routes -> server
	config.InitConfig()
	fmt.Println("eu estava chorando até agoraaa (igor guimaraes)")
	database := database.NewDatabase()
	repository := repository.NewRepositories(database.DB)
	useCase := usecase.NewUseCase(repository)
	controller := controller.NewController(useCase)

	server := server.NewServer(controller)
	server.PrepareRoutes()

	// TODO - safe shutdown
	ctx := context.Background()
	server.Listen(ctx, config.CFG.SERVER_HOST, config.CFG.SERVER_PORT)

}
