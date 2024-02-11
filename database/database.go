package database

import (
	"context"
	"fmt"
	"log"

	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-demode/config"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	// DB *sql.DB
	DB *mongo.Client
}

func NewDatabase() *Database {

	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.CFG.DB_USER, config.CFG.DB_PASSWORD, config.CFG.DB_HOST, config.CFG.DB_PORT)
	clientOpts := options.Client().ApplyURI(url)
	clientOpts = clientOpts.SetMaxPoolSize(config.CFG.DB_MAX_POOL_SIZE)
	clientOpts = clientOpts.SetMinPoolSize(config.CFG.DB_MIN_POOL_SIZE)
	clientOpts = clientOpts.SetMaxConnecting(config.CFG.DB_MAX_CONNECTIONS)
	clientOpts = clientOpts.SetCompressors([]string{config.CFG.DB_COMPRESSORS})
	clientOpts = clientOpts.SetZlibLevel(config.CFG.DB_ZLIB_COMPRESS_LEVEL)
	clientOpts = clientOpts.SetZstdLevel(config.CFG.DB_ZSTD_COMPRESS_LEVEL)

	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	_ = client
	return &Database{DB: client}

}
