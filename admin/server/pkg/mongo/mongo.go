package mongo

import (
	"bytes"
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"time"
)

func init() {
	// Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	viper.SetDefault("MONGODB_URI", "mongodb://127.0.0.1:27017")
	viper.SetDefault("MONGODB_DATABASE", "fsme")
	viper.SetDefault("MONGODB_COLLECTION", "fsm")
	viper.SetDefault("MONGODB_SINGLE", false)

	// Set config from .ENV
	if _, err := os.Stat(".env"); err == nil {
		config, err := ioutil.ReadFile(".env")
		if err != nil {
			logger.Error("Error read .env file", zap.Error(err))
		}

		viper.SetConfigType("env")
		viper.ReadConfig(bytes.NewBuffer(config))
	}
}

var Cfg DB

func Run() (*DB, error) {
	var err error

	// Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	mongoCfg := options.Client().ApplyURI(viper.GetString("MONGODB_URI"))
	mongoCfg.SetDirect(viper.GetBool("MONGODB_SINGLE")) // for work with single node
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Cfg.Client, err = mongo.Connect(ctx, mongoCfg)

	if err != nil {
		return &Cfg, err
	}

	// Check the connection
	ctxping, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = Cfg.Client.Ping(ctxping, readpref.Primary())

	if err != nil {
		return &Cfg, err
	}

	logger.Info("Run mongodb")

	// Setting
	Cfg.DB = Cfg.Client.Database(viper.GetString("MONGODB_DATABASE"))

	return &Cfg, err
}

func (db *DB) Close() {
	db.Client.Disconnect(db.Context)
}
