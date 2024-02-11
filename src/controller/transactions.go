package controller

import (
	"encoding/json"
	"net/http"
	"slices"
	"strconv"

	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-demode/src/interfaces"
	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-demode/src/usecase"
	"github.com/gin-gonic/gin"
)

type (
	TransactionsController interface {
		CreateTransaction(context *gin.Context)
		ConsultaExtrato(context *gin.Context)
	}

	transactions struct {
		useCases usecase.TransactionsUseCase
	}
)

func NewTransactionController(uc usecase.TransactionsUseCase) TransactionsController {
	return &transactions{useCases: uc}
}

func (t *transactions) ConsultaExtrato(context *gin.Context) {
	id, canGetId := context.Params.Get("id")

	if id == "" || !canGetId {
		context.JSON(http.StatusNotFound, "Cliente não encontrado")
		return
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusNotFound, "Cliente não encontrado")
		return
	}

	result, status := t.useCases.ConsultaExtrato(idInt)
	if status != 200 {
		context.JSON(status, result)
		return
	}
	context.JSON(status, result)

}

func (t *transactions) CreateTransaction(context *gin.Context) {
	id, canGetId := context.Params.Get("id")

	if id == "" || !canGetId {
		context.JSON(http.StatusNotFound, "Cliente não encontrado")
		return
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusNotFound, "Cliente não encontrado")
		return
	}

	request := &interfaces.RequestCreateTransaction{}

	err = json.NewDecoder(context.Request.Body).Decode(request)

	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	request.ID = idInt

	operacao := []string{"c", "d"}
	if !slices.Contains(operacao, request.Tipo) {
		context.JSON(http.StatusUnprocessableEntity, "tipo de operação invalida")
		return
	}

	if request.Valor <= 0 {
		context.JSON(http.StatusUnprocessableEntity, "valor invalido")
		return
	}

	if !(len(request.Descricao) >= 1 && len(request.Descricao) <= 10) {
		context.JSON(http.StatusUnprocessableEntity, "descricao invalida")
		return
	}

	// go func(ww http.ResponseWriter, rr *http.Request) {
	result, status := t.useCases.CreateTransaction(*request)
	if status != 200 {
		context.JSON(status, result)
		return
	}

	context.JSON(200, result)
	// }(w, r)
}
