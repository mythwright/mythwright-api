package db

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const mappingCollection = "item_mappings"

type ItemMapping struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	AccountID  primitive.ObjectID `bson:"account_id,omitempty"`
	ItemID     int                `bson:"item_id" json:"item_id,omitempty"`
	ChildrenID []int              `bson:"children_id" json:"children_id,omitempty"`
}

func (d *Database) SaveItemMappings(ctx context.Context, im ItemMapping) error {
	db := d.db.Database("mythwright")
	col := db.Collection("mappings")
	_, err := col.InsertOne(ctx, im)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("unable to save doc")
		return err
	}
	return nil
}

func (d *Database) GetItemMappings(ctx context.Context, acctID primitive.ObjectID, itemID int) (*ItemMapping, error) {
	col := d.DefaultDBCollection(mappingCollection)
	im := &ItemMapping{}
	if err := col.FindOne(ctx, bson.D{{"account_id", acctID}, {"item_id", itemID}}).Decode(im); err != nil {
		return nil, err
	}
	return im, nil
}
