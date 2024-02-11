package usecase

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-demode/src/interfaces"
	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-demode/src/repository"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type (
	TransactionsUseCase interface {
		CreateTransaction(req interfaces.RequestCreateTransaction) (*interfaces.ResponseCreateTransaction, int)
		ConsultaExtrato(cliente int64) (*interfaces.ResponseGetExtrato, int)
	}

	transactions struct {
		repo repository.TransactionsRepository
	}
)

func NewTransactions(repo repository.TransactionsRepository) TransactionsUseCase {
	return &transactions{repo: repo}
}

// ConsultaSaldo implements Transactions.
func (t *transactions) ConsultaExtrato(cliente int64) (*interfaces.ResponseGetExtrato, int) {

	extratos, err := t.repo.ConsultaExtrato(nil, cliente, 10)
	if err != nil {
		log.Println("ERRORRR 1")
		log.Println(err.Error())
		log.Println(err)
		return nil, 500
	}
	if len(extratos) == 0 {
		return nil, 404
	}

	result := &interfaces.ResponseGetExtrato{}
	result.Saldo.Total = extratos[0].Saldo
	result.Saldo.Limite = extratos[0].Limite
	result.Saldo.DataExtrato = time.Now().UTC().Format("2006-01-02T15:04:05-0700")

	for _, ext := range extratos {

		newTran := interfaces.UltimasTransacoes{}
		if ext.TipoTransacao != nil {
			newTran.Descricao = *ext.Descricao
			newTran.RealizadaEm = ext.CreatedAt
			newTran.Tipo = *ext.TipoTransacao
			newTran.Valor = *ext.Valor
			result.UltimasTransacoes = append(result.UltimasTransacoes, newTran)
		}
	}

	return result, http.StatusOK
}

// CreateTransaction implements Transactions.
func (t *transactions) CreateTransaction(req interfaces.RequestCreateTransaction) (*interfaces.ResponseCreateTransaction, int) {
	wc := writeconcern.Majority()
	txnOptions := options.Transaction().SetWriteConcern(wc)

	session, err := t.repo.CriaSessionDb()
	if err != nil {
		log.Println("ERRORRR 0 - SESSION")
		log.Println(err.Error())
		log.Println(err)
		return nil, 500
	}
	defer session.EndSession(context.TODO())

	session.StartTransaction(txnOptions)
	session.Client()

	extratos, err := t.repo.ConsultaExtrato(session.Client(), req.ID, 1)
	if err != nil {
		log.Println("ERRORRR 2")
		log.Println(err.Error())
		log.Println(err)
		session.AbortTransaction(context.TODO())
		return nil, 500
	}

	if len(extratos) == 0 {
		session.AbortTransaction(context.TODO())
		return nil, 404
	}

	if req.Tipo == "c" {
		extratos[0].Saldo = extratos[0].Saldo + req.Valor
	} else {
		extratos[0].Saldo = extratos[0].Saldo - req.Valor
		if extratos[0].Saldo+extratos[0].Limite < 0 {
			session.AbortTransaction(context.TODO())
			return nil, 422
		}
	}

	extratos[0].TipoTransacao = &req.Tipo
	extratos[0].Descricao = &req.Descricao
	extratos[0].Valor = &req.Valor

	err = t.repo.CriaTransacao(session.Client(), extratos[0])
	if err != nil {
		log.Println("ERRORRR 3")
		log.Println(err.Error())
		session.AbortTransaction(context.TODO())
		return nil, 500
	}

	return &interfaces.ResponseCreateTransaction{Limite: extratos[0].Limite, Saldo: extratos[0].Saldo}, http.StatusOK
}
