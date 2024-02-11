package server

import (
	"context"

	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-demode/src/controller"
	"github.com/gin-gonic/gin"
)

type (
	Server interface {
		PrepareRoutes()
		Listen(ctx context.Context, url string, port string)
	}
	server struct {
		transaction controller.TransactionsController
		gin         *gin.Engine
	}
)

func NewServer(controller controller.Controller) Server {
	return &server{gin: gin.Default(), transaction: controller.NewTransactionController()}
}

func (s *server) PrepareRoutes() {
	root := s.gin.Group("")
	root.GET("clientes/:id/extrato", s.transaction.ConsultaExtrato)
	root.POST("clientes/:id/transacoes", s.transaction.CreateTransaction)

}

func (s *server) Listen(ctx context.Context, url string, port string) {
	// s.gin.Run(fmt.Sprintf(":%s", port))
	err := s.gin.RunTLS(":8085", "./server-cert.pem", "./server.key")
	if err != nil {
		panic(err)
	}

}
