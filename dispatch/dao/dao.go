package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hk4e/common/config"
	"hk4e/logger"
)

type Dao struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewDao() (r *Dao) {
	r = new(Dao)
	clientOptions := options.Client().ApplyURI(config.CONF.Database.Url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.LOG.Error("mongo connect error: %v", err)
		return nil
	}
	r.client = client
	r.db = client.Database("gate_hk4e")
	return r
}

func (d *Dao) CloseDao() {
	err := d.client.Disconnect(context.TODO())
	if err != nil {
		logger.LOG.Error("mongo close error: %v", err)
	}
}
