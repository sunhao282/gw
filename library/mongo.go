package library

import (
	"context"
	"fmt"
	"time"

	"gw/backend"
	"gw/conf"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//链接mongodb
var cli *mongo.Database

func init() {
	host := fmt.Sprintf("mongodb://%s:%s", conf.MongoDB["host"], conf.MongoDB["port"])
	opts := &options.ClientOptions{}
	opts.SetMaxPoolSize(conf.MongoPoll)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	con, err := mongo.Connect(ctx, options.Client().ApplyURI(host), opts)
	if err != nil {
		panic("mongodb connect error")
	}

	cli = con.Database(conf.MongoDB["dbname"])
}

//写入tb记录
func Add(tb string, m interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := cli.Collection(tb).InsertOne(ctx, m); err != nil {
		return err
	}

	return nil
}

//查询一条数据
func FindOne(tb string, w bson.M, m interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := cli.Collection(tb).FindOne(ctx, w).Decode(m); err != nil {
		return err
	}
	return nil
}

//查询全部数据
func FindAll(tb string) ([]backend.MongoInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := cli.Collection(tb).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	if err := cur.Err(); err != nil {
		return nil, err
	}

	var info []backend.MongoInfo
	for cur.Next(ctx) {
		var dec *backend.MongoInfo
		if err := cur.Decode(&dec); err != nil {
			return nil, err
		}
		info = append(info, *dec)
	}

	return info, nil
}

//删除
func Del(tb string, w bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := cli.Collection(tb).DeleteOne(ctx, w); err != nil {
		return err
	}

	return nil
}
