package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pwd  string `yaml:"pwd"`
	DB   string `yaml:"db"`
}

var Mongo *mongo.Database

func InitMongo(c *MongoConfig) {
	var url = fmt.Sprintf("mongodb://%s:%s", c.Host, c.Port)
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Panicf("new mongo client err:%v", err)
	}
	var ctx = context.TODO()
	err = client.Connect(ctx)
	if err != nil {
		log.Panicf("conn mongo err:%v", err)
	}
	Mongo = client.Database(c.DB)
}

type Table interface {
	TableName() string
}
type NTable interface {
	TableName(v int64) string
}

func GetMongo(v interface{}, x ...int64) *mongo.Collection {
	if len(x) > 0 {
		if i, ok := v.(NTable); ok {
			return Mongo.Collection(i.TableName(x[0]))
		}
	}
	if i, ok := v.(Table); ok {
		return Mongo.Collection(i.TableName())
	}

	return nil
}
