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
