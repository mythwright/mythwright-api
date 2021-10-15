package db

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	db *mongo.Client
}

func NewDatabase(ctx context.Context) (*Database, error) {
	d := &Database{}

	client, err := mongo.Connect(ctx, createOpts())
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: (%v)", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("unable to ping to database: (%v)", err)
	}

	return d, nil
}

func createOpts() *options.ClientOptions {
	opts := options.Client()
	opts.ApplyURI(viper.GetString("db_uri"))
	return opts
}

func (d *Database) DefaultDBCollection(col string, opts ...*options.CollectionOptions) *mongo.Collection {
	return d.db.Database(viper.GetString("default_db")).Collection(col, opts...)
}
