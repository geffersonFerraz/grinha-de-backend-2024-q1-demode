package repository

import (
	"context"

	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-demode/src/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	TransactionsRepository interface {
		CriaTransacao(mongo *mongo.Client, req interfaces.ExtratoFromDB) error
		ConsultaExtrato(mongo *mongo.Client, client int64, limit int64) ([]interfaces.ExtratoFromDB, error)
		CriaSessionDb() (mongo.Session, error)
	}

	transactions struct {
		DB *mongo.Client
	}
)

func NewTransactions(db *mongo.Client) TransactionsRepository {
	return &transactions{DB: db}
}

func (t *transactions) ConsultaExtrato(mongo *mongo.Client, client int64, limit int64) ([]interfaces.ExtratoFromDB, error) {
	if mongo == nil {
		mongo = t.DB
	}
	filter := bson.D{{Key: "id_cliente", Value: client}}
	coll := mongo.Database("rinha").Collection("clientes")

	findOptions := options.Find()
	// 1 = ASC	// -1 = DESC
	findOptions.SetSort(bson.D{{"_id", -1}})
	findOptions.SetLimit(limit)

	cursor, err := coll.Find(context.TODO(), filter, findOptions)

	if err != nil {
		return nil, err
	}

	var results []interfaces.ExtratoFromDB
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (t *transactions) CriaTransacao(mongo *mongo.Client, req interfaces.ExtratoFromDB) error {

	_, err := mongo.Database("rinha").Collection("clientes").InsertOne(context.TODO(), req)

	if err != nil {
		return err
	}

	return nil
}

func (t *transactions) CriaSessionDb() (mongo.Session, error) {

	return t.DB.StartSession()

}
